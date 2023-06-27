// storage.go
package storage

import (
	"fmt"
	"io"
	"os"

	"github.com/bwoff11/frens/internal/config"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/google/uuid"
)

type Storage interface {
	SaveFile(data []byte, filename string) error
	LoadFile(id uuid.UUID) (io.ReadCloser, error)
	DeleteFile(id uuid.UUID) error
}

func NewStorage(cfg *config.Config) (Storage, error) {
	storageConfig := cfg.Storage
	switch storageConfig.Type {
	case "local":
		logger.Log.Info().Msg("Local storage initialized")
		localStorage := newLocalStorage(&storageConfig)

		// Check if the directory exists, and if not, create it
		if _, err := os.Stat(localStorage.Path); os.IsNotExist(err) {
			os.MkdirAll(localStorage.Path, 0755)
		}

		return localStorage, nil
	case "s3":
		logger.Log.Info().Msg("S3 storage initialized")
		// Similar directory checks could be made here for your S3 bucket
		return newS3Storage(&storageConfig), nil
	default:
		logger.Log.Error().Msg("Unsupported storage type: " + storageConfig.Type)
		return nil, fmt.Errorf("unsupported storage type: %s", storageConfig.Type)
	}
}
