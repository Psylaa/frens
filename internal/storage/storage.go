package storage

import (
	"fmt"
	"io"

	"github.com/bwoff11/frens/internal/config"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/google/uuid"
)

type Storage interface {
	DetectFileType(data []byte) config.FileType
	SaveFile(id uuid.UUID, data []byte) error
	LoadFile(id uuid.UUID) (io.ReadCloser, error)
	DeleteFile(id uuid.UUID) error
}

// storage.go
func NewStorage(cfg *config.Config) (map[config.FileType]Storage, error) {
	storages := make(map[config.FileType]Storage)
	for fileType, storageConfig := range cfg.Storage {
		switch storageConfig.Type {
		case "local":
			storages[fileType] = newLocalStorage(&storageConfig)
			logger.Log.Info().Msg("Local storage for " + string(fileType) + " initialized")
		case "s3":
			storages[fileType] = newS3Storage(&storageConfig)
			logger.Log.Info().Msg("S3 storage for " + string(fileType) + " initialized")
		default:
			logger.Log.Error().Msg("Unsupported storage type: " + storageConfig.Type)
			return nil, fmt.Errorf("unsupported storage type: %s", storageConfig.Type)
		}
	}

	logger.Log.Info().Msg("Storage initialized")
	return storages, nil
}
