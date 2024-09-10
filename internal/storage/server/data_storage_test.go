package server

import (
	"context"
	"errors"
	"testing"

	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	logger "github.com/Stern-Ritter/gophkeeper/internal/logger/server"
	"github.com/Stern-Ritter/gophkeeper/internal/model"
)

func TestDataStorageCreate_Success(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err, "Error init connection mock")
	defer mock.Close()

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	dataStorage := NewDataStorage(mock, l)

	data := model.Data{
		UserID:        "1",
		Type:          model.AccountType,
		SensitiveData: []byte("sensitive data"),
		Comment:       "comment",
	}

	mock.ExpectExec(`
		INSERT INTO gophkeeper\.data \(user_id, type, data, comment\)
		VALUES \(@userID, @type, @data, @comment\)
	`).
		WithArgs(
			data.UserID,
			string(data.Type),
			data.SensitiveData,
			data.Comment).
		WillReturnResult(pgxmock.NewResult("INSERT", 1))

	err = dataStorage.Create(context.Background(), data)
	assert.NoError(t, err, "Error inserting data")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "The expected sql commands were not executed")
}

func TestDataStorageCreate_Error(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err, "Error init connection mock")
	defer mock.Close()

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	dataStorage := NewDataStorage(mock, l)

	data := model.Data{
		UserID:        "1",
		Type:          model.AccountType,
		SensitiveData: []byte("sensitive data"),
		Comment:       "comment",
	}

	mock.ExpectExec(`
		INSERT INTO gophkeeper\.data \(user_id, type, data, comment\)
		VALUES \(@userID, @type, @data, @comment\)
	`).
		WithArgs(
			data.UserID,
			string(data.Type),
			data.SensitiveData,
			data.Comment).
		WillReturnError(errors.New("database error"))

	err = dataStorage.Create(context.Background(), data)
	assert.Error(t, err, "Expected an error when inserting data")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "The expected sql commands were not executed")
}

func TestDataStorageDelete_Success(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err, "Error init connection mock")
	defer mock.Close()

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	dataStorage := NewDataStorage(mock, l)

	dataID := "1"

	mock.ExpectExec(`
		DELETE FROM gophkeeper\.data WHERE id = @id
	`).
		WithArgs(dataID).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	err = dataStorage.Delete(context.Background(), dataID)
	assert.NoError(t, err, "Error deleting data")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "The expected sql commands were not executed")
}

func TestDataStorageDelete_Error(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err, "Error init connection mock")
	defer mock.Close()

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	dataStorage := NewDataStorage(mock, l)

	dataID := "1"

	mock.ExpectExec(`
		DELETE FROM gophkeeper\.data WHERE id = @id
	`).
		WithArgs(dataID).
		WillReturnError(errors.New("database error"))

	err = dataStorage.Delete(context.Background(), dataID)
	assert.Error(t, err, "Expected an error when deleting data")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "The expected sql commands were not executed")
}

func TestDataStorageGetByID_Success(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err, "Error init connection mock")
	defer mock.Close()

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	dataStorage := NewDataStorage(mock, l)

	dataID := "1"
	expectedData := model.Data{
		ID:            dataID,
		UserID:        "1",
		SensitiveData: []byte("sensitive data"),
		Type:          model.AccountType,
		Comment:       "comment",
	}

	rows := pgxmock.NewRows([]string{"id", "user_id", "data", "type", "comment"}).
		AddRow(expectedData.ID, expectedData.UserID, expectedData.SensitiveData, string(expectedData.Type), expectedData.Comment)

	mock.ExpectQuery(`
		SELECT id, user_id, data, type, comment FROM gophkeeper\.data WHERE id = @id
	`).
		WithArgs(dataID).
		WillReturnRows(rows)

	data, err := dataStorage.GetByID(context.Background(), dataID)
	assert.NoError(t, err, "Error getting data by ID")
	assert.Equal(t, expectedData, data, "Returned data does not match expected")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "The expected sql commands were not executed")
}

func TestDataStorageGetByID_Error(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err, "Error init connection mock")
	defer mock.Close()

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	dataStorage := NewDataStorage(mock, l)

	dataID := "1"

	mock.ExpectQuery(`
		SELECT id, user_id, data, type, comment FROM gophkeeper\.data WHERE id = @id
	`).
		WithArgs(dataID).
		WillReturnError(errors.New("database error"))

	data, err := dataStorage.GetByID(context.Background(), dataID)
	assert.Error(t, err, "Expected an error when getting data by ID")
	assert.Equal(t, model.Data{}, data, "Returned data should be empty on error")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "The expected sql commands were not executed")
}

func TestDataStorageGetAll_Success(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err, "Error init connection mock")
	defer mock.Close()

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	dataStorage := NewDataStorage(mock, l)

	userID := "1"
	expectedData := []model.Data{
		{
			ID:            "1",
			UserID:        userID,
			SensitiveData: []byte("encrypted data 1"),
			Type:          model.CardType,
			Comment:       "comment 1",
		},
		{
			ID:            "2",
			UserID:        userID,
			SensitiveData: []byte("encrypted data 2"),
			Type:          model.CardType,
			Comment:       "comment 2",
		},
	}

	rows := pgxmock.NewRows([]string{"id", "user_id", "data", "comment"}).
		AddRow(expectedData[0].ID, expectedData[0].UserID, expectedData[0].SensitiveData, expectedData[0].Comment).
		AddRow(expectedData[1].ID, expectedData[1].UserID, expectedData[1].SensitiveData, expectedData[1].Comment)

	mock.ExpectQuery(`
		SELECT id, user_id, data, comment FROM gophkeeper\.data WHERE user_id = @userID AND type = @type
	`).
		WithArgs(userID, string(model.CardType)).
		WillReturnRows(rows)

	data, err := dataStorage.GetAll(context.Background(), userID, model.CardType)
	assert.NoError(t, err, "Error getting all data")
	assert.Equal(t, expectedData, data, "Returned data does not match expected")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "The expected sql commands were not executed")
}

func TestDataStorageGetAll_Error(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err, "Error init connection mock")
	defer mock.Close()

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	dataStorage := NewDataStorage(mock, l)

	userID := "1"

	mock.ExpectQuery(`
		SELECT id, user_id, data, comment FROM gophkeeper\.data WHERE user_id = @userID AND type = @type
	`).
		WithArgs(userID, string(model.CardType)).
		WillReturnError(errors.New("database error"))

	data, err := dataStorage.GetAll(context.Background(), userID, model.CardType)
	assert.Error(t, err, "Expected an error when getting all data")
	assert.Nil(t, data, "Returned data should be nil on error")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "The expected sql commands were not executed")
}
