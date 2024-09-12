package server

import (
	"context"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	er "github.com/Stern-Ritter/gophkeeper/internal/errors"
	logger "github.com/Stern-Ritter/gophkeeper/internal/logger/server"
	"github.com/Stern-Ritter/gophkeeper/internal/model"
	pb "github.com/Stern-Ritter/gophkeeper/proto/gen/gophkeeper/gophkeeperapi/v1"
)

func TestUploadFile_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFileStorage := NewMockFileStorage(ctrl)
	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	fileService := NewFileService(mockFileStorage, "/tmp", l)

	stream := NewMockFileServiceV1_UploadFileServer(ctrl)

	userID := "1"
	fileName := "test.txt"
	comment := "test comment"
	fileData := []byte("file data")

	stream.EXPECT().Recv().Return(&pb.UploadFileRequestV1{
		Name:    fileName,
		Comment: comment,
	}, nil)

	stream.EXPECT().Recv().Return(&pb.UploadFileRequestV1{
		Data: fileData,
	}, nil).Times(1)

	stream.EXPECT().Recv().Return(nil, io.EOF)

	mockFileStorage.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)

	err = fileService.UploadFile(context.Background(), userID, stream)
	assert.NoError(t, err, "Expected no error during file upload")
}

func TestUploadFile_RecvMetadataError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFileStorage := NewMockFileStorage(ctrl)
	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	fileService := NewFileService(mockFileStorage, "/tmp", l)

	stream := NewMockFileServiceV1_UploadFileServer(ctrl)
	stream.EXPECT().Recv().Return(nil, fmt.Errorf("failed to receive data"))

	err = fileService.UploadFile(context.Background(), "1", stream)
	assert.Error(t, err, "Expected error due to stream receiving failure")
}

func TestUploadFile_RecvFilePartError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFileStorage := NewMockFileStorage(ctrl)
	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	fileService := NewFileService(mockFileStorage, "/tmp", l)

	stream := NewMockFileServiceV1_UploadFileServer(ctrl)

	userID := "1"
	fileName := "test.txt"
	comment := "test comment"

	stream.EXPECT().Recv().Return(&pb.UploadFileRequestV1{
		Name:    fileName,
		Comment: comment,
	}, nil)

	stream.EXPECT().Recv().Return(nil, fmt.Errorf("failed to receive data"))

	err = fileService.UploadFile(context.Background(), userID, stream)
	assert.Error(t, err, "Expected error due to stream receiving failure")
}

func TestUploadFile_InternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFileStorage := NewMockFileStorage(ctrl)
	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	fileService := NewFileService(mockFileStorage, "/tmp", l)

	stream := NewMockFileServiceV1_UploadFileServer(ctrl)

	userID := "1"
	fileName := "test.txt"
	comment := "test comment"
	fileData := []byte("file data")

	stream.EXPECT().Recv().Return(&pb.UploadFileRequestV1{
		Name:    fileName,
		Comment: comment,
	}, nil)

	stream.EXPECT().Recv().Return(&pb.UploadFileRequestV1{
		Data: fileData,
	}, nil).Times(1)

	stream.EXPECT().Recv().Return(nil, io.EOF)

	mockFileStorage.EXPECT().Create(gomock.Any(), gomock.Any()).Return(fmt.Errorf("internal error"))

	err = fileService.UploadFile(context.Background(), userID, stream)
	assert.Error(t, err, "Expected internal error during file upload")
}

func TestDownloadFile_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFileStorage := NewMockFileStorage(ctrl)
	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	fileService := NewFileService(mockFileStorage, "/tmp", l)

	stream := NewMockFileServiceV1_DownloadFileServer(ctrl)

	userID := "1"
	fileID := "1"
	filePath := "/tmp/test.txt"
	fileData := []byte("file data")

	fileMetadata := model.File{
		ID:     fileID,
		UserID: userID,
		Name:   "test.txt",
		Size:   int64(len(fileData)),
		Path:   filePath,
	}

	mockFileStorage.EXPECT().GetByID(gomock.Any(), fileID).Return(fileMetadata, nil)

	stream.EXPECT().Send(&pb.DownloadFileResponseV1{
		Name: fileMetadata.Name,
		Size: fileMetadata.Size,
	}).Return(nil)

	stream.EXPECT().Send(gomock.Any()).DoAndReturn(func(resp *pb.DownloadFileResponseV1) error {
		assert.Equal(t, fileData, resp.Data, "Expected the file data to be sent")
		return nil
	}).Times(1)

	tmpFile, err := os.Create(filePath)
	require.NoError(t, err)
	defer tmpFile.Close()
	defer os.Remove(filePath)
	_, err = tmpFile.Write(fileData)
	require.NoError(t, err)

	err = fileService.DownloadFile(context.Background(), userID, fileID, stream)
	assert.NoError(t, err, "Expected no error during file download")
}

func TestDownloadFile_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFileStorage := NewMockFileStorage(ctrl)
	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	fileService := NewFileService(mockFileStorage, "/tmp", l)

	stream := NewMockFileServiceV1_DownloadFileServer(ctrl)

	userID := "1"
	fileID := "1"

	mockFileStorage.EXPECT().GetByID(gomock.Any(), fileID).Return(model.File{}, pgx.ErrNoRows)

	err = fileService.DownloadFile(context.Background(), userID, fileID, stream)
	assert.ErrorAs(t, err, &er.NotFoundError{}, "Expected NotFoundError when file is not found")
}

func TestDownloadFile_SendMetadataError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFileStorage := NewMockFileStorage(ctrl)
	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	fileService := NewFileService(mockFileStorage, "/tmp", l)

	stream := NewMockFileServiceV1_DownloadFileServer(ctrl)

	userID := "1"
	fileID := "1"
	filePath := "/tmp/test.txt"
	fileData := []byte("file data")

	fileMetadata := model.File{
		ID:     fileID,
		UserID: userID,
		Name:   "test.txt",
		Size:   int64(len(fileData)),
		Path:   filePath,
	}

	mockFileStorage.EXPECT().GetByID(gomock.Any(), fileID).Return(fileMetadata, nil)

	stream.EXPECT().Send(&pb.DownloadFileResponseV1{
		Name: fileMetadata.Name,
		Size: fileMetadata.Size,
	}).Return(fmt.Errorf("send metadata error"))

	tmpFile, err := os.Create(filePath)
	require.NoError(t, err, "unexpected error creating tmp file")
	defer tmpFile.Close()
	defer os.Remove(filePath)
	_, err = tmpFile.Write(fileData)
	require.NoError(t, err, "unexpected error writing to tmp file")

	err = fileService.DownloadFile(context.Background(), userID, fileID, stream)
	assert.Error(t, err, "Expected error when sending file metadata fails")
}

func TestFileServiceImplDownloadFile_SendFilePartError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFileStorage := NewMockFileStorage(ctrl)
	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	fileService := NewFileService(mockFileStorage, "/tmp", l)

	stream := NewMockFileServiceV1_DownloadFileServer(ctrl)

	userID := "1"
	fileID := "1"
	filePath := "/tmp/test.txt"
	fileData := []byte("file data")

	fileMetadata := model.File{
		ID:     fileID,
		UserID: userID,
		Name:   "test.txt",
		Size:   int64(len(fileData)),
		Path:   filePath,
	}

	mockFileStorage.EXPECT().GetByID(gomock.Any(), fileID).Return(fileMetadata, nil)

	stream.EXPECT().Send(&pb.DownloadFileResponseV1{
		Name: fileMetadata.Name,
		Size: fileMetadata.Size,
	}).Return(nil)

	stream.EXPECT().Send(gomock.Any()).Return(fmt.Errorf("send file part error")).Times(1)

	tmpFile, err := os.Create(filePath)
	require.NoError(t, err, "unexpected error creating tmp file")
	defer tmpFile.Close()
	defer os.Remove(filePath)
	_, err = tmpFile.Write(fileData)
	require.NoError(t, err, "unexpected error writing to tmp file")

	err = fileService.DownloadFile(context.Background(), userID, fileID, stream)
	assert.Error(t, err, "Expected error when sending file part fails")
}

func TestDownloadFile_InternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFileStorage := NewMockFileStorage(ctrl)
	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	fileService := NewFileService(mockFileStorage, "/tmp", l)

	stream := NewMockFileServiceV1_DownloadFileServer(ctrl)

	userID := "1"
	fileID := "1"

	mockFileStorage.EXPECT().GetByID(gomock.Any(), fileID).Return(model.File{}, fmt.Errorf("internal error"))

	err = fileService.DownloadFile(context.Background(), userID, fileID, stream)
	assert.Error(t, err, "Expected internal error during file download")
}

func TestDeleteFile_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFileStorage := NewMockFileStorage(ctrl)
	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	fileService := NewFileService(mockFileStorage, "/tmp", l)

	userID := "1"
	fileID := "1"
	filePath := "/tmp/test.txt"

	fileMetadata := model.File{
		ID:     fileID,
		UserID: userID,
		Path:   filePath,
	}

	mockFileStorage.EXPECT().GetByID(gomock.Any(), fileID).Return(fileMetadata, nil)

	mockFileStorage.EXPECT().Delete(gomock.Any(), fileID).Return(nil)

	tmpFile, err := os.Create(filePath)
	require.NoError(t, err)
	defer tmpFile.Close()
	defer os.Remove(filePath)

	err = fileService.DeleteFile(context.Background(), userID, fileID)
	assert.NoError(t, err, "Expected no error during file deletion")
}

func TestDeleteFile_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFileStorage := NewMockFileStorage(ctrl)
	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	fileService := NewFileService(mockFileStorage, "/tmp", l)

	userID := "1"
	fileID := "1"

	mockFileStorage.EXPECT().GetByID(gomock.Any(), fileID).Return(model.File{}, pgx.ErrNoRows)

	err = fileService.DeleteFile(context.Background(), userID, fileID)
	assert.ErrorAs(t, err, &er.NotFoundError{}, "Expected NotFoundError when file is not found")
}

func TestDeleteFile_AccessDenied(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFileStorage := NewMockFileStorage(ctrl)
	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	fileService := NewFileService(mockFileStorage, "/tmp", l)

	userID := "1"
	fileID := "1"
	filePath := "/tmp/test.txt"

	fileMetadata := model.File{
		ID:     fileID,
		UserID: "2",
		Path:   filePath,
	}

	mockFileStorage.EXPECT().GetByID(gomock.Any(), fileID).Return(fileMetadata, nil)

	err = fileService.DeleteFile(context.Background(), userID, fileID)
	assert.ErrorAs(t, err, &er.ForbiddenError{},
		"Expected ForbiddenError when user attempts to delete file not belonging to them")
}

func TestDeleteFile_GetByIDInternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFileStorage := NewMockFileStorage(ctrl)
	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	fileService := NewFileService(mockFileStorage, "/tmp", l)

	userID := "1"
	fileID := "1"

	mockFileStorage.EXPECT().GetByID(gomock.Any(), fileID).Return(model.File{}, fmt.Errorf("internal error"))

	err = fileService.DeleteFile(context.Background(), userID, fileID)
	assert.Error(t, err, "Expected internal error when retrieving file metadata for deletion")
}

func TestDeleteFile_DeleteInternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFileStorage := NewMockFileStorage(ctrl)
	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	fileService := NewFileService(mockFileStorage, "/tmp", l)

	userID := "1"
	fileID := "1"
	filePath := "/tmp/test.txt"

	fileMetadata := model.File{
		ID:     fileID,
		UserID: userID,
		Path:   filePath,
	}

	mockFileStorage.EXPECT().GetByID(gomock.Any(), fileID).Return(fileMetadata, nil)
	mockFileStorage.EXPECT().Delete(gomock.Any(), fileID).Return(fmt.Errorf("internal error"))

	tmpFile, err := os.Create(filePath)
	require.NoError(t, err)
	defer tmpFile.Close()
	defer os.Remove(filePath)

	err = fileService.DeleteFile(context.Background(), userID, fileID)
	assert.Error(t, err, "Expected internal error during file deletion")
}

func TestGetFileByID_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFileStorage := NewMockFileStorage(ctrl)
	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	fileService := NewFileService(mockFileStorage, "/tmp", l)

	userID := "1"
	fileID := "1"

	expectedFile := model.File{
		ID:     fileID,
		UserID: userID,
		Name:   "test.txt",
		Size:   1024,
		Path:   "/tmp/test.txt",
	}

	mockFileStorage.EXPECT().GetByID(gomock.Any(), fileID).Return(expectedFile, nil)

	file, err := fileService.GetFileByID(context.Background(), userID, fileID)
	assert.NoError(t, err, "Expected no error retrieving file by ID")
	assert.Equal(t, expectedFile, file, "Expected file to match")
}

func TestGetFileByID_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFileStorage := NewMockFileStorage(ctrl)
	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	fileService := NewFileService(mockFileStorage, "/tmp", l)

	userID := "1"
	fileID := "1"

	mockFileStorage.EXPECT().GetByID(gomock.Any(), fileID).Return(model.File{}, pgx.ErrNoRows)

	_, err = fileService.GetFileByID(context.Background(), userID, fileID)
	assert.ErrorAs(t, err, &er.NotFoundError{}, "Expected NotFoundError when file is not found")
}

func TestGetFileByID_InternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFileStorage := NewMockFileStorage(ctrl)
	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	fileService := NewFileService(mockFileStorage, "/tmp", l)

	userID := "1"
	fileID := "1"

	mockFileStorage.EXPECT().GetByID(gomock.Any(), fileID).Return(model.File{}, fmt.Errorf("internal error"))

	_, err = fileService.GetFileByID(context.Background(), userID, fileID)
	assert.Error(t, err, "Expected internal error during GetFileByID")
}

func TestGetAllFiles_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFileStorage := NewMockFileStorage(ctrl)
	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	fileService := NewFileService(mockFileStorage, "/tmp", l)

	userID := "1"
	expectedFiles := []model.File{
		{ID: "1", UserID: userID, Name: "file1.txt", Size: 1024, Path: "/tmp/file1.txt"},
		{ID: "2", UserID: userID, Name: "file2.pdf", Size: 2048, Path: "/tmp/file2.pdf"},
	}

	mockFileStorage.EXPECT().GetAll(gomock.Any(), userID).Return(expectedFiles, nil)

	files, err := fileService.GetAllFiles(context.Background(), userID)
	assert.NoError(t, err, "Expected no error getting all files")
	assert.Equal(t, expectedFiles, files, "Expected files list to match")
}

func TestGetAllFiles_InternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFileStorage := NewMockFileStorage(ctrl)
	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	fileService := NewFileService(mockFileStorage, "/tmp", l)

	userID := "1"

	mockFileStorage.EXPECT().GetAll(gomock.Any(), userID).Return([]model.File{}, fmt.Errorf("internal error"))

	files, err := fileService.GetAllFiles(context.Background(), userID)
	assert.Error(t, err, "Expected internal error during GetAllFiles")
	assert.Empty(t, files, "Expected no files data due to internal error")
}
