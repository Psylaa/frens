package storage

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/bwoff11/frens/internal/config"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/google/uuid"
	"github.com/h2non/filetype"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type FileType string

const (
	Image FileType = "image"
	Video FileType = "video"
	Audio FileType = "audio"
	Other FileType = "other"
)

func InitStorage(cfg *config.Config) error {
	logger.Log.Info().Msg("Initializing storage")

	for _, storageConfig := range cfg.Storage {
		if storageConfig.Type == "local" {
			if _, err := os.Stat(storageConfig.Local.Path); os.IsNotExist(err) {
				err := os.MkdirAll(storageConfig.Local.Path, 0755)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func DetectFileType(data []byte) FileType {
	kind, err := filetype.Match(data)
	if err != nil {
		return Other
	}

	switch kind.MIME.Type {
	case "image":
		return Image
	case "video":
		return Video
	case "audio":
		return Audio
	default:
		return Other
	}
}

func SaveFile(id uuid.UUID, data []byte) error {
	fileType := DetectFileType(data)
	storageConfig, ok := config.Config.Storage[fileType]
	if !ok {
		return fmt.Errorf("unsupported file type: %s", fileType)
	}

	return saveToStorage(id, data, storageConfig)
}

func LoadFile(id uuid.UUID) (io.ReadCloser, error) {
	fileType := DetectFileType(nil)
	storageConfig, ok := config.Config.Storage[fileType]
	if !ok {
		return nil, fmt.Errorf("unsupported file type: %s", fileType)
	}

	data, err := loadFromStorage(id, nil, storageConfig)
	if err != nil {
		return nil, err
	}

	return ioutil.NopCloser(bytes.NewReader(data)), nil
}

func DeleteFile(id uuid.UUID) error {
	fileType := DetectFileType(nil)
	storageConfig, ok := config.Config.Storage[fileType]
	if !ok {
		return fmt.Errorf("unsupported file type: %s", fileType)
	}

	return deleteFromStorage(id, nil, storageConfig)
}

func performStorageOperation(id uuid.UUID, data []byte, operation func(id uuid.UUID, data []byte, cfg config.StorageDetails) error) error {
	fileType := DetectFileType(data)
	storageConfig, ok := config.Config.Storage[fileType]
	if !ok {
		return fmt.Errorf("unsupported file type: %s", fileType)
	}

	return operation(id, data, storageConfig)
}

func saveToStorage(id uuid.UUID, data []byte, cfg config.StorageDetails) error {
	switch cfg.Type {
	case "local":
		fullPath := filepath.Join(cfg.Local.Path, id.String())
		return ioutil.WriteFile(fullPath, data, 0644)
	case "s3":
		sess := session.Must(session.NewSession(&aws.Config{
			Region: aws.String(cfg.S3.Region),
		}))

		uploader := s3manager.NewUploader(sess)
		_, err := uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(cfg.S3.Bucket),
			Key:    aws.String(id.String()),
			Body:   bytes.NewReader(data),
		})

		return err
	default:
		return fmt.Errorf("unsupported storage type: %s", cfg.Type)
	}
}

func loadFromStorage(id uuid.UUID, data []byte, cfg config.StorageDetails) ([]byte, error) {
	switch cfg.Type {
	case "local":
		fullPath := filepath.Join(cfg.Local.Path, id.String())
		return ioutil.ReadFile(fullPath)
	case "s3":
		sess := session.Must(session.NewSession(&aws.Config{
			Region: aws.String(cfg.S3.Region),
		}))

		downloader := s3manager.NewDownloader(sess)

		buf := aws.NewWriteAtBuffer([]byte{})
		_, err := downloader.Download(buf, &s3.GetObjectInput{
			Bucket: aws.String(cfg.S3.Bucket),
			Key:    aws.String(id.String()),
		})

		return buf.Bytes(), err
	default:
		return nil, fmt.Errorf("unsupported storage type: %s", cfg.Type)
	}
}

func deleteFromStorage(id uuid.UUID, data []byte, cfg config.StorageDetails) error {
	switch cfg.Type {
	case "local":
		fullPath := filepath.Join(cfg.Local.Path, id.String())
		return os.Remove(fullPath)
	case "s3":
		sess := session.Must(session.NewSession(&aws.Config{
			Region: aws.String(cfg.S3.Region),
		}))

		svc := s3.New(sess)
		_, err := svc.DeleteObject(&s3.DeleteObjectInput{
			Bucket: aws.String(cfg.S3.Bucket),
			Key:    aws.String(id.String()),
		})

		return err
	default:
		return fmt.Errorf("unsupported storage type: %s", cfg.Type)
	}
}
