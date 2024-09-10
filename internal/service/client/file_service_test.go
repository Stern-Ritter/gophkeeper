package client

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	pb "github.com/Stern-Ritter/gophkeeper/proto/gen/gophkeeper/gophkeeperapi/v1"
)

func TestUploadFile(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockClient := NewMockFileServiceV1Client(mockCtrl)
	fileService := NewFileService(mockClient)

	filePath := "test.txt"
	comment := "comment"
	fileData := []byte("test data")
	progressFunc := func(progress float64) {}

	tmpFile, err := os.Create(filePath)
	require.NoError(t, err, "Expected no error creating temp file")
	defer os.Remove(filePath)
	defer tmpFile.Close()

	_, err = tmpFile.Write(fileData)
	require.NoError(t, err, "Expected no error writing test data to temp file")

	fileInfo, err := tmpFile.Stat()
	require.NoError(t, err, "Expected no error getting file stat info")

	stream := NewMockFileServiceV1_UploadFileClient(mockCtrl)

	mockClient.EXPECT().UploadFile(gomock.Any()).Return(stream, nil)

	stream.EXPECT().Send(&pb.UploadFileRequestV1{
		Name:    filepath.Base(filePath),
		Comment: comment,
	}).Return(nil)

	stream.EXPECT().Send(gomock.Any()).
		DoAndReturn(func(req *pb.UploadFileRequestV1) error {
			assert.Equal(t, int64(len(req.GetData())), fileInfo.Size(), "Expected the full file data to be sent")
			return nil
		}).Times(1)

	stream.EXPECT().CloseAndRecv().Return(&pb.UploadFileResponseV1{}, nil)

	err = fileService.UploadFile(context.Background(), filePath, comment, progressFunc)
	assert.NoError(t, err, "Expected no error uploading file")
}

func TestDownloadFile(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockClient := NewMockFileServiceV1Client(mockCtrl)
	fileService := NewFileService(mockClient)

	fileID := "1"
	dirPath := "."
	fileName := "test.txt"
	fileData := []byte("test data")
	progressFunc := func(progress float64) {}

	stream := NewMockFileServiceV1_DownloadFileClient(mockCtrl)
	mockClient.EXPECT().DownloadFile(gomock.Any(), &pb.DownloadFileRequestV1{Id: fileID}).Return(stream, nil)

	stream.EXPECT().Recv().Return(&pb.DownloadFileResponseV1{
		Name: fileName,
		Size: int64(len(fileData)),
	}, nil)

	stream.EXPECT().Recv().DoAndReturn(func() (*pb.DownloadFileResponseV1, error) {
		return &pb.DownloadFileResponseV1{Data: fileData}, nil
	})

	stream.EXPECT().Recv().Return(nil, io.EOF)

	err := fileService.DownloadFile(context.Background(), fileID, dirPath, progressFunc)
	assert.NoError(t, err, "Expected no error downloading file")

	filePath := filepath.Join(dirPath, fileName)
	downloadedData, err := os.ReadFile(filePath)
	defer os.Remove(filePath)

	assert.NoError(t, err, "Expected no error reading downloaded file")
	assert.Equal(t, fileData, downloadedData, "Expected downloaded data to match sended")
}

func TestGetAllFiles(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockClient := NewMockFileServiceV1Client(mockCtrl)
	fileService := NewFileService(mockClient)

	expectedFiles := []*pb.FileV1{
		{Id: "1", Name: "file1.txt", Size: 1024, Comment: "comment1"},
		{Id: "2", Name: "file2.pdf", Size: 2048, Comment: "comment2"},
	}

	mockClient.EXPECT().GetFiles(gomock.Any(), &pb.GetFilesRequestV1{}).Return(&pb.GetFilesResponseV1{
		Files: expectedFiles,
	}, nil)

	files, err := fileService.GetAllFiles()
	assert.NoError(t, err, "Expected no error getting files")
	assert.Equal(t, expectedFiles, files, "Expected files list to match")
}

func TestDeleteFile(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockClient := NewMockFileServiceV1Client(mockCtrl)
	fileService := NewFileService(mockClient)

	fileID := "1"

	mockClient.EXPECT().DeleteFile(gomock.Any(), &pb.DeleteFileRequestV1{Id: fileID}).Return(&pb.DeleteFileResponseV1{}, nil)

	err := fileService.DeleteFile(fileID)
	assert.NoError(t, err, "Expected no error deleting file")
}
