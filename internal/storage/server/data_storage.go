package server

import (
	"context"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"

	logger "github.com/Stern-Ritter/gophkeeper/internal/logger/server"
	"github.com/Stern-Ritter/gophkeeper/internal/model"
)

// DataStorage defines the interface for data storage operations.
type DataStorage interface {
	Create(ctx context.Context, data model.Data) error
	Delete(ctx context.Context, dataID string) error
	GetByID(ctx context.Context, dataID string) (model.Data, error)
	GetAll(ctx context.Context, userID string, dataType model.DataType) ([]model.Data, error)
}

// DataStorageImpl is an implementation of the DataStorage interface.
type DataStorageImpl struct {
	db     PgxIface
	Logger *logger.ServerLogger
}

// NewDataStorage creates a new instance of DataStorage.
func NewDataStorage(db PgxIface, logger *logger.ServerLogger) DataStorage {
	return &DataStorageImpl{
		db:     db,
		Logger: logger,
	}
}

// Create inserts a new data entry into the database.
func (s *DataStorageImpl) Create(ctx context.Context, data model.Data) error {
	_, err := s.db.Exec(ctx, `
		INSERT INTO gophkeeper.data (user_id, type, data, comment)
		VALUES (@userID, @type, @data, @comment)
	`, pgx.NamedArgs{
		"userID":  data.UserID,
		"type":    string(data.Type),
		"data":    data.SensitiveData,
		"comment": data.Comment,
	})

	if err != nil {
		s.Logger.Debug("Failed to insert data", zap.String("event", "add data"),
			zap.String("user id", data.UserID), zap.String("data type", string(data.Type)), zap.Error(err))
		return err
	}

	return nil
}

// Delete removes a data entry from the database by its ID.
func (s *DataStorageImpl) Delete(ctx context.Context, dataID string) error {
	_, err := s.db.Exec(ctx, `
		DELETE FROM gophkeeper.data WHERE id = @id
	`, pgx.NamedArgs{
		"id": dataID,
	})

	if err != nil {
		s.Logger.Debug("Failed to delete data", zap.String("event", "delete data"),
			zap.String("data id", dataID), zap.Error(err))
		return err
	}

	return nil
}

// GetByID retrieves a data entry from the database by its ID.
func (s *DataStorageImpl) GetByID(ctx context.Context, dataID string) (model.Data, error) {
	row := s.db.QueryRow(ctx, `
	SELECT id, user_id, data, type, comment
	FROM gophkeeper.data 
	WHERE id = @id
	`, pgx.NamedArgs{
		"id": dataID,
	})

	data := model.Data{}
	var dataType string
	err := row.Scan(&data.ID, &data.UserID, &data.SensitiveData, &dataType, &data.Comment)

	if err != nil {
		s.Logger.Debug("Failed to get data", zap.String("event", "get data"),
			zap.String("data id", dataID), zap.Error(err))
		return data, err
	}
	data.Type = model.DataType(dataType)

	return data, nil
}

// GetAll retrieves all data entries for a specific user and data type from the database.
func (s *DataStorageImpl) GetAll(ctx context.Context, userID string, dataType model.DataType) ([]model.Data, error) {
	rows, err := s.db.Query(ctx, `
		SELECT id, user_id, data, comment
		FROM gophkeeper.data
		WHERE user_id = @userID AND type = @type
	`, pgx.NamedArgs{
		"userID": userID,
		"type":   string(dataType),
	})

	if err != nil {
		s.Logger.Debug("Failed to select all data", zap.String("event", "get all data"),
			zap.String("user id", userID), zap.String("data type", string(dataType)), zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	allData := make([]model.Data, 0)

	for rows.Next() {
		data := model.Data{
			Type: dataType,
		}

		if err = rows.Scan(&data.ID, &data.UserID, &data.SensitiveData, &data.Comment); err != nil {
			s.Logger.Debug("Failed to select all data", zap.String("event", "get all data"),
				zap.String("user id", userID), zap.String("data type", string(dataType)), zap.Error(err))
			return nil, err
		}

		allData = append(allData, data)
	}

	if err = rows.Err(); err != nil {
		s.Logger.Debug("Failed to select all data", zap.String("event", "get all data"),
			zap.String("user id", userID), zap.String("data type", string(dataType)), zap.Error(err))
		return nil, err
	}

	return allData, nil
}
