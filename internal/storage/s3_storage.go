package storage

import (
	"io"

	"github.com/bwoff11/frens/internal/config"
	"github.com/google/uuid"
)

type S3Storage struct {
	Bucket string
	Region string
}

func newS3Storage(cfg *config.StorageDetails) *S3Storage {
	return &S3Storage{
		Bucket: cfg.S3.Bucket,
		Region: cfg.S3.Region,
	}
}

func (s3s *S3Storage) SaveFile(data []byte, filename string) error {
	return nil
}

func (s3s *S3Storage) LoadFile(id uuid.UUID) (io.ReadCloser, error) {
	return nil, nil
}

func (s3s *S3Storage) DeleteFile(id uuid.UUID) error {
	return nil
}
