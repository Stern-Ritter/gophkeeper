package server

import (
	"context"
	"fmt"
	"testing"

	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	logger "github.com/Stern-Ritter/gophkeeper/internal/logger/server"
	"github.com/Stern-Ritter/gophkeeper/internal/model"
)

func TestFileStorageCreate_Success(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err, "Error init connection mock")
	defer mock.Close()

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	fileStorage := NewFileStorage(mock, l)

	file := model.File{
		UserID:  "1",
		Name:    "file.txt",
		Size:    1024,
		Path:    "/storage/file.txt",
		Comment: "comment",
	}

	mock.ExpectExec(`
		INSERT INTO gophkeeper.files \(user_id, name, size, path, comment\)
		VALUES \(@user_id, @name, @size, @path, @comment\)
	`).
		WithArgs(file.UserID, file.Name, file.Size, file.Path, file.Comment).
		WillReturnResult(pgxmock.NewResult("INSERT", 1))

	err = fileStorage.Create(context.Background(), file)
	assert.NoError(t, err, "Error creating file metadata")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "The expected sql commands were not executed")
}

func TestFileStorageCreate_Error(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err, "Error init connection mock")
	defer mock.Close()

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	fileStorage := NewFileStorage(mock, l)

	file := model.File{
		UserID:  "1",
		Name:    "file.txt",
		Size:    1024,
		Path:    "/storage/file.txt",
		Comment: "comment",
	}

	mock.ExpectExec(`
		INSERT INTO gophkeeper.files \(user_id, name, size, path, comment\)
		VALUES \(@user_id, @name, @size, @path, @comment\)
	`).
		WithArgs(file.UserID, file.Name, file.Size, file.Path, file.Comment).
		WillReturnError(fmt.Errorf("insert error"))

	err = fileStorage.Create(context.Background(), file)
	assert.Error(t, err, "Expected error when creating file metadata")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "The expected sql commands were not executed")
}

func TestFileStorageDelete_Success(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err, "Error init connection mock")
	defer mock.Close()

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	fileStorage := NewFileStorage(mock, l)

	fileID := "1"

	mock.ExpectExec(`
		DELETE FROM gophkeeper.files WHERE id = @id
	`).
		WithArgs(fileID).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	err = fileStorage.Delete(context.Background(), fileID)
	assert.NoError(t, err, "Error deleting file metadata")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "The expected sql commands were not executed")
}

func TestFileStorageDelete_Error(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err, "Error init connection mock")
	defer mock.Close()

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	fileStorage := NewFileStorage(mock, l)

	fileID := "1"

	mock.ExpectExec(`
		DELETE FROM gophkeeper.files WHERE id = @id
	`).
		WithArgs(fileID).
		WillReturnError(fmt.Errorf("delete error"))

	err = fileStorage.Delete(context.Background(), fileID)
	assert.Error(t, err, "Expected error when deleting file metadata")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "The expected sql commands were not executed")
}

func TestFileStorageGetByID_Success(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err, "Error init connection mock")
	defer mock.Close()

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	fileStorage := NewFileStorage(mock, l)

	fileID := "1"
	expectedFile := model.File{
		ID:      fileID,
		UserID:  "1",
		Name:    "file.txt",
		Size:    1024,
		Path:    "/storage/file.txt",
		Comment: "comment",
	}

	row := pgxmock.NewRows([]string{"id", "user_id", "name", "size", "path", "comment"}).
		AddRow(expectedFile.ID, expectedFile.UserID, expectedFile.Name, expectedFile.Size, expectedFile.Path, expectedFile.Comment)

	mock.ExpectQuery(`
		SELECT id, user_id, name, size, path, comment
		FROM gophkeeper.files
		WHERE id = @id
	`).
		WithArgs(fileID).
		WillReturnRows(row)

	file, err := fileStorage.GetByID(context.Background(), fileID)
	assert.NoError(t, err, "Error getting file metadata by ID")
	assert.Equal(t, expectedFile, file, "Returned file metadata does not match expected")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "The expected sql commands were not executed")
}

func TestFileStorageGetByID_Error(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err, "Error init connection mock")
	defer mock.Close()

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	fileStorage := NewFileStorage(mock, l)

	fileID := "1"

	mock.ExpectQuery(`
		SELECT id, user_id, name, size, path, comment
		FROM gophkeeper.files
		WHERE id = @id
	`).
		WithArgs(fileID).
		WillReturnError(fmt.Errorf("query error"))

	file, err := fileStorage.GetByID(context.Background(), fileID)
	assert.Error(t, err, "Expected error when getting file metadata by ID")
	assert.Equal(t, model.File{}, file, "Returned file metadata should be empty on error")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "The expected sql commands were not executed")
}

func TestFileStorageGetAll_Success(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err, "Error init connection mock")
	defer mock.Close()

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	fileStorage := NewFileStorage(mock, l)

	userID := "1"
	expectedFiles := []model.File{
		{
			ID:      "1",
			UserID:  userID,
			Name:    "file1.txt",
			Size:    1024,
			Path:    "/storage/file1.txt",
			Comment: "comment 1",
		},
		{
			ID:      "2",
			UserID:  userID,
			Name:    "file2.txt",
			Size:    2048,
			Path:    "/storage/file2.txt",
			Comment: "comment 2",
		},
	}

	rows := pgxmock.NewRows([]string{"id", "user_id", "name", "size", "path", "comment"})
	for _, file := range expectedFiles {
		rows.AddRow(file.ID, file.UserID, file.Name, file.Size, file.Path, file.Comment)
	}

	mock.ExpectQuery(`
		SELECT id, user_id, name, size, path, comment
		FROM gophkeeper.files
		WHERE user_id = @userId
	`).
		WithArgs(userID).
		WillReturnRows(rows)

	files, err := fileStorage.GetAll(context.Background(), userID)
	assert.NoError(t, err, "Error getting all file metadata by user ID")
	assert.Equal(t, expectedFiles, files, "Returned files metadata does not match expected")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "The expected sql commands were not executed")
}

func TestFileStorageGetAll_Error(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err, "Error init connection mock")
	defer mock.Close()

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	fileStorage := NewFileStorage(mock, l)

	userID := "1"

	mock.ExpectQuery(`
		SELECT id, user_id, name, size, path, comment
		FROM gophkeeper.files
		WHERE user_id = @userId
	`).
		WithArgs(userID).
		WillReturnError(fmt.Errorf("query error"))

	files, err := fileStorage.GetAll(context.Background(), userID)
	assert.Error(t, err, "Expected error when getting all file metadata")
	assert.Nil(t, files, "Returned files metadata should be nil on error")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "The expected sql commands were not executed")
}
