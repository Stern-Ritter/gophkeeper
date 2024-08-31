package server

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"

	er "github.com/Stern-Ritter/gophkeeper/internal/errors"
	logger "github.com/Stern-Ritter/gophkeeper/internal/logger/server"
	storage "github.com/Stern-Ritter/gophkeeper/internal/storage/server"
	pb "github.com/Stern-Ritter/gophkeeper/proto/gen/gophkeeper/gophkeeperapi/v1"

	"github.com/Stern-Ritter/gophkeeper/internal/model"
)

const (
	filePartSize = 1024 * 1024 // filePartSize defines the size of file parts used during file upload and download operations.
)

// FileService defines the interface for file operations such as uploading, downloading,
// deleting, and retrieving files.
type FileService interface {
	UploadFile(ctx context.Context, userID string, stream pb.FileServiceV1_UploadFileServer) error
	DownloadFile(ctx context.Context, userID string, fileID string, stream pb.FileServiceV1_DownloadFileServer) error
	DeleteFile(ctx context.Context, userID string, fileID string) error
	GetFileByID(ctx context.Context, userID string, fileID string) (model.File, error)
	GetAllFiles(ctx context.Context, userID string) ([]model.File, error)
}

// FileServiceImpl provides a concrete implementation of the FileService interface.
type FileServiceImpl struct {
	fileStorage     storage.FileStorage
	fileStoragePath string
	logger          *logger.ServerLogger
}

// NewFileService creates a new instance of FileService.
func NewFileService(fileStorage storage.FileStorage, fileStoragePath string, logger *logger.ServerLogger) FileService {
	return &FileServiceImpl{
		fileStorage:     fileStorage,
		fileStoragePath: fileStoragePath,
		logger:          logger,
	}
}

// UploadFile handles uploading a file from the client with gRPC stream.
// It creates a new file in the user's storage directory and writes file parts as they are received.
func (s *FileServiceImpl) UploadFile(ctx context.Context, userID string, stream pb.FileServiceV1_UploadFileServer) (err error) {
	req, err := stream.Recv()
	if err != nil {
		s.logger.Error("Failed to get file metadata", zap.String("event", "upload file"),
			zap.String("user id", userID), zap.Error(err))
		return err
	}

	fileName := req.GetName()
	comment := req.GetComment()

	filesDir := filepath.Join(s.fileStoragePath, userID)
	filePath := filepath.Join(filesDir, fileName)
	if _, err = os.Stat(filesDir); os.IsNotExist(err) {
		err = os.MkdirAll(filesDir, 0755)
		if err != nil {
			s.logger.Error("Failed to create user file storage directory",
				zap.String("event", "upload file"), zap.String("user id", userID), zap.Error(err))
			return err
		}
	}

	file, err := os.Create(filePath)
	if err != nil {
		s.logger.Error("Failed to create file in user file storage directory",
			zap.String("event", "upload file"), zap.String("user id", userID), zap.Error(err))
		return err
	}
	defer file.Close()

	var filePart []byte
	var fileSize int64
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				s.logger.Error("Failed to get file part", zap.String("event", "upload file"),
					zap.String("user id", userID), zap.Error(err))
				removeFileDueError(filePath, userID, err, s.logger)
				return err
			}
		}

		filePart = req.GetData()
		n, err := file.Write(filePart)
		if err != nil {
			s.logger.Error("Failed to write file part to file", zap.String("event", "upload file"),
				zap.String("user id", userID), zap.Error(err))
			removeFileDueError(filePath, userID, err, s.logger)
			return err
		}
		fileSize += int64(n)
	}

	fileData := model.File{
		UserID:  userID,
		Name:    fileName,
		Size:    fileSize,
		Path:    filePath,
		Comment: comment,
	}

	err = s.fileStorage.Create(ctx, fileData)
	if err != nil {
		s.logger.Error("Failed to save file metadata to database", zap.String("event", "upload file"),
			zap.String("user id", userID), zap.Error(err))
		removeFileDueError(filePath, userID, err, s.logger)
		return err
	}

	return nil
}

func removeFileDueError(filePath string, userID string, err error, logger *logger.ServerLogger) {
	logger.Error("Removing file due to error", zap.String("file", filePath),
		zap.String("user id", userID), zap.Error(err))
	removeFileErr := os.Remove(filePath)
	if err != nil {
		logger.Error("Failed to remove file", zap.String("file", filePath),
			zap.String("user id", userID), zap.Error(removeFileErr))
	}
}

// DownloadFile handles downloading a file to the client with gRPC stream.
// It reads file parts from the file in the user's storage directory and sends them to the user.
func (s *FileServiceImpl) DownloadFile(ctx context.Context, userID string, fileID string, stream pb.FileServiceV1_DownloadFileServer) error {
	fileMetadata, err := s.fileStorage.GetByID(ctx, fileID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			s.logger.Error("Failed to get file metadata from database", zap.String("file id", fileID),
				zap.String("user id", userID), zap.Error(err))
			return er.NewNotFoundError(fmt.Sprintf("file with id: %s not found", fileID), err)
		}
	}

	if fileMetadata.UserID != userID {
		s.logger.Warn("User attempted to access file belonging to another user", zap.String("file id", fileID),
			zap.String("user id", userID), zap.Error(err))
		return er.NewForbiddenError("access denied", err)
	}

	file, err := os.Open(fileMetadata.Path)
	if err != nil {
		s.logger.Error("Failed to open file for downloading", zap.String("file path", fileMetadata.Path),
			zap.String("user id", userID), zap.Error(err))
		return err
	}
	defer file.Close()

	downloadFileMetaData := &pb.DownloadFileResponseV1{
		Name: fileMetadata.Name,
		Size: fileMetadata.Size,
	}
	if err = stream.Send(downloadFileMetaData); err != nil {
		s.logger.Error("Failed to send download file metadata", zap.String("file path", fileMetadata.Path),
			zap.String("user id", userID), zap.Error(err))
		return err
	}

	buffer := make([]byte, filePartSize)
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				s.logger.Error("Failed to read file part", zap.String("file path", fileMetadata.Path),
					zap.String("user id", userID), zap.Error(err))
				return err
			}
		}

		filePart := &pb.DownloadFileResponseV1{
			Data: buffer[:n],
		}

		if err = stream.Send(filePart); err != nil {
			s.logger.Error("Failed to send file part", zap.String("file path", fileMetadata.Path),
				zap.String("user id", userID), zap.Error(err))
			return err
		}
	}

	return nil
}

// DeleteFile deletes a file identified by file ID belonging to a user.
// It removes the file from the database and the file from user storage directory.
func (s *FileServiceImpl) DeleteFile(ctx context.Context, userID string, fileID string) error {
	file, err := s.fileStorage.GetByID(ctx, fileID)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			s.logger.Info("File does not exist", zap.String("event", "delete file"),
				zap.String("file id", fileID), zap.String("user id", userID), zap.Error(err))
			return er.NewNotFoundError("file does not exist", err)
		default:
			s.logger.Error("Failed to get file by id", zap.String("event", "delete file"),
				zap.String("file id", fileID), zap.String("user id", userID), zap.Error(err))
			return fmt.Errorf("failed to get file by id: %w", err)
		}
	}

	if file.UserID != userID {
		s.logger.Warn("User attempted to access file belonging to another user",
			zap.String("event", "delete file"), zap.String("file id", fileID),
			zap.String("user id", userID), zap.Error(err))
		return er.NewForbiddenError("access denied", err)
	}

	err = s.fileStorage.Delete(ctx, file.ID)
	if err != nil {
		s.logger.Error("Failed to delete file from database", zap.String("event", "delete file"),
			zap.String("file id", file.ID), zap.String("user id", userID), zap.Error(err))
		return err
	}

	err = os.Remove(file.Path)
	if err != nil {
		s.logger.Error("Failed to delete file from user file storage",
			zap.String("event", "delete file"), zap.String("file path", file.Path),
			zap.String("user id", userID), zap.Error(err))
		return err
	}

	return nil
}

// GetFileByID get a file's metadata by file ID.
func (s *FileServiceImpl) GetFileByID(ctx context.Context, userID string, fileID string) (model.File, error) {
	file, err := s.fileStorage.GetByID(ctx, fileID)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			s.logger.Info("File does not exist", zap.String("event", "get file by id"),
				zap.String("file id", fileID), zap.String("user id", userID), zap.Error(err))
			return model.File{}, er.NewNotFoundError("file does not exist", err)
		default:
			s.logger.Error("Failed to get file by id", zap.String("event", "get file by id"),
				zap.String("file id", fileID), zap.String("user id", userID), zap.Error(err))
			return model.File{}, fmt.Errorf("failed to get file by id: %w", err)
		}
	}

	return file, nil
}

// GetAllFiles get all files metadata belonging to a specific user.
func (s *FileServiceImpl) GetAllFiles(ctx context.Context, userID string) ([]model.File, error) {
	files, err := s.fileStorage.GetAll(ctx, userID)
	if err != nil {
		s.logger.Error("Failed to get files", zap.String("event", "get files"),
			zap.String("user id", userID), zap.Error(err))
		return []model.File{}, fmt.Errorf("failed to get files: %w", err)
	}

	return files, nil
}
