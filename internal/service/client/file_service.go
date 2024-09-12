package client

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	pb "github.com/Stern-Ritter/gophkeeper/proto/gen/gophkeeper/gophkeeperapi/v1"
)

const (
	filePartSize = 1024 * 1024 // filePartSize defines the size of each file part that is sent or received during file upload and download operations.
)

// FileService defines an interface for managing file operations such as uploading, downloading,
// retrieving, and deleting files.
type FileService interface {
	UploadFile(ctx context.Context, filePath string, comment string, progressFunc func(float64)) error
	DownloadFile(ctx context.Context, fileID string, dirPath string, progressFunc func(float64)) error
	GetAllFiles() ([]*pb.FileV1, error)
	DeleteFile(fileID string) error
}

// FileServiceImpl is a concrete implementation of the FileService interface.
type FileServiceImpl struct {
	fileClient pb.FileServiceV1Client
}

// NewFileService creates a new instance of FileService.
func NewFileService(fileClient pb.FileServiceV1Client) FileService {
	return &FileServiceImpl{
		fileClient: fileClient,
	}
}

// UploadFile uploads a file to the server by reading it parts defined by filePartSize.
// The progress of the upload is reported with the progressFunc callback.
func (s *FileServiceImpl) UploadFile(ctx context.Context, filePath string, comment string, progressFunc func(float64)) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("file '%s' does not exist or is opened by another application", filePath)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file info: %w", err)
	}

	stream, err := s.fileClient.UploadFile(ctx)
	if err != nil {
		return fmt.Errorf("failed to open file send stream: %w", err)
	}

	name := filepath.Base(filePath)
	fileMetadata := &pb.UploadFileRequestV1{
		Name:    name,
		Comment: comment,
	}
	if err := stream.Send(fileMetadata); err != nil {
		return fmt.Errorf("failed to send file metadata: %w", err)
	}

	totalSize := fileInfo.Size()
	totalUpload := 0
	buffer := make([]byte, filePartSize)
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return fmt.Errorf("failed to read file part: %w", err)
			}
		}

		filePart := &pb.UploadFileRequestV1{
			Data: buffer[:n],
		}
		if err := stream.Send(filePart); err != nil {
			return fmt.Errorf("failed to send file part: %w", err)
		}

		totalUpload += n

		progress := float64(totalUpload) / float64(totalSize) * 100
		if progressFunc != nil {
			progressFunc(progress)
		}
	}

	_, err = stream.CloseAndRecv()
	if err != nil {
		return fmt.Errorf("failed to receive succes upload response: %w", err)
	}

	return nil
}

// DownloadFile downloads a file from the server by receiving it in parts defined by filePartSize.
// The progress of the download is reported with the progressFunc callback.
func (s *FileServiceImpl) DownloadFile(ctx context.Context, fileID string, dirPath string, progressFunc func(float64)) error {
	stream, err := s.fileClient.DownloadFile(ctx, &pb.DownloadFileRequestV1{Id: fileID})
	if err != nil {
		return fmt.Errorf("failed to open file download stream: %w", err)
	}

	resp, err := stream.Recv()
	if err != nil {
		return fmt.Errorf("failed to get file metadata: %w", err)
	}

	fileName := resp.GetName()
	filePath := filepath.Join(dirPath, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	totalSize := resp.GetSize()
	totalDownload := 0
	for {
		resp, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				os.Remove(filePath)
				return fmt.Errorf("failed to get file part: %w", err)
			}
		}

		filePart := resp.GetData()
		n, err := file.Write(filePart)
		if err != nil {
			os.Remove(filePath)
			return fmt.Errorf("failed to write file part to file: %w", err)
		}

		totalDownload += n
		progress := float64(totalDownload) / float64(totalSize) * 100
		if progressFunc != nil {
			progressFunc(progress)
		}
	}

	return nil
}

// GetAllFiles get a list of all files stored on the server by sending a request to the file service with gRPC.
func (s *FileServiceImpl) GetAllFiles() ([]*pb.FileV1, error) {
	ctx := context.Background()
	req := &pb.GetFilesRequestV1{}
	resp, err := s.fileClient.GetFiles(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.Files, nil
}

// DeleteFile deletes a file from the server identified by the given file id by sending a request
// to the file service with gRPC. Returns an error if the file deletion fails.
func (s *FileServiceImpl) DeleteFile(fileID string) error {
	ctx := context.Background()
	req := &pb.DeleteFileRequestV1{Id: fileID}

	_, err := s.fileClient.DeleteFile(ctx, req)
	return err
}
