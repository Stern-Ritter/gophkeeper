package server

import (
	"context"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"

	logger "github.com/Stern-Ritter/gophkeeper/internal/logger/server"
	"github.com/Stern-Ritter/gophkeeper/internal/model"
)

// FileStorage defines the interface for file metadata storage operations.
type FileStorage interface {
	Create(ctx context.Context, file model.File) error
	Delete(ctx context.Context, fileID string) error
	GetByID(ctx context.Context, fileID string) (model.File, error)
	GetAll(ctx context.Context, userID string) ([]model.File, error)
}

// FileStorageImpl is an implementation of the FileStorage interface.
type FileStorageImpl struct {
	db     PgxIface
	Logger logger.ServerLogger
}

// NewFileStorage creates a new instance of FileStorage.
func NewFileStorage(db PgxIface, logger logger.ServerLogger) FileStorage {
	return &FileStorageImpl{
		db:     db,
		Logger: logger,
	}
}

// Create inserts a new file metadata entry into the database.
func (s *FileStorageImpl) Create(ctx context.Context, file model.File) error {
	_, err := s.db.Exec(ctx, `
		INSERT INTO gophkeeper.files (user_id, name, size, path, comment)
		VALUES (@user_id, @name, @size, @path, @comment)
	`, pgx.NamedArgs{
		"user_id": file.UserID,
		"name":    file.Name,
		"size":    file.Size,
		"path":    file.Path,
		"comment": file.Comment,
	})

	if err != nil {
		s.Logger.Debug("Failed to insert file metadata", zap.String("event", "add file metadata"),
			zap.Error(err))
	}

	return err
}

// Delete removes a file metadata entry from the database by its ID.
func (s *FileStorageImpl) Delete(ctx context.Context, fileID string) error {
	_, err := s.db.Exec(ctx, `
		DELETE FROM gophkeeper.files WHERE id = @id
	`, pgx.NamedArgs{"id": fileID})

	if err != nil {
		s.Logger.Debug("Failed to delete file metadata", zap.String("event", "delete file metadata"),
			zap.String("file id", fileID), zap.Error(err))
	}

	return err
}

// GetByID retrieves a file metadata entry from the database by its ID.
func (s *FileStorageImpl) GetByID(ctx context.Context, fileID string) (model.File, error) {
	file := model.File{}

	row := s.db.QueryRow(ctx, `
		SELECT id, user_id, name, size, path, comment
		FROM gophkeeper.files
		WHERE id = @id
	`, pgx.NamedArgs{
		"id": fileID,
	})

	err := row.Scan(&file.ID, &file.UserID, &file.Name, &file.Size, &file.Path, &file.Comment)
	if err != nil {
		s.Logger.Debug("Failed to get file metadata", zap.String("event", "get file metadata"),
			zap.String("file id", fileID), zap.Error(err))
		return file, err
	}

	return file, nil
}

// GetAll get all file metadata entries for a specific user from the database.
func (s *FileStorageImpl) GetAll(ctx context.Context, userID string) ([]model.File, error) {
	rows, err := s.db.Query(ctx, `
		SELECT id, user_id, name, size, path, comment
		FROM gophkeeper.files
		WHERE user_id = @userId
	`, pgx.NamedArgs{
		"userId": userID,
	})

	if err != nil {
		s.Logger.Debug("Failed to get all user files metadata", zap.String("event", "get all files metadata"),
			zap.String("user id", userID), zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	files := make([]model.File, 0)

	for rows.Next() {
		file := model.File{}

		if err = rows.Scan(&file.ID, &file.UserID, &file.Name, &file.Size, &file.Path, &file.Comment); err != nil {
			s.Logger.Debug("Failed to get all files metadata", zap.String("event", "get all files metadata"),
				zap.Error(err))
			return nil, err
		}

		files = append(files, file)
	}

	if err = rows.Err(); err != nil {
		s.Logger.Debug("Failed to get all files metadata", zap.String("event", "get all files metadata"),
			zap.Error(err))
		return nil, err
	}

	return files, nil
}
