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

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type FileType string

const (
	FileTypeImage       FileType = "profile_picture"
	FileTypeVideo       FileType = "user_banner"
	FileTypeStatusFile  FileType = "status_file"
	FileTypeStatusImage FileType = "status_image"
	FileTypeStatusVideo FileType = "status_video"
	FileTypeStatusAudio FileType = "status_audio"
)

var storageConfigs map[FileType]config.StorageDetails

func InitStorage(cfg *config.Config) error {
	logger.Log.Info().Msg("Initializing storage")

	// Store storage configs in map for easy access
	storageConfigs = map[FileType]config.StorageDetails{
		FileTypeImage:       cfg.Storage.ProfilePictures,
		FileTypeVideo:       cfg.Storage.UserBanners,
		FileTypeStatusFile:  cfg.Storage.StatusFiles,
		FileTypeStatusImage: cfg.Storage.StatusImages,
		FileTypeStatusVideo: cfg.Storage.StatusVideos,
		FileTypeStatusAudio: cfg.Storage.StatusAudio,
	}

	// Create local directories if the storage type is local
	for _, storageConfig := range storageConfigs {
		if storageConfig.Type == "local" {
			dir := storageConfig.Local.Path
			if _, err := os.Stat(dir); os.IsNotExist(err) {
				os.MkdirAll(dir, 0755)
			}
		}
	}

	return nil
}

func SaveFile(fileType FileType, path string, data []byte) error {
	config, ok := storageConfigs[fileType]
	if !ok {
		return fmt.Errorf("invalid file type: %s", fileType)
	}

	switch config.Type {
	case "local":
		fullPath := filepath.Join(config.Local.Path, path)
		return os.WriteFile(fullPath, data, 0644)
	case "s3":
		sess := session.Must(session.NewSession(&aws.Config{
			Region: aws.String(config.S3.Region),
		}))

		uploader := s3manager.NewUploader(sess)

		_, err := uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(config.S3.Bucket),
			Key:    aws.String(path),
			Body:   bytes.NewReader(data),
		})

		return err
	default:
		return fmt.Errorf("unsupported storage type: %s", config.Type)
	}
}

func LoadFile(fileType FileType, path string) (io.ReadCloser, error) {
	config, ok := storageConfigs[fileType]
	if !ok {
		return nil, fmt.Errorf("invalid file type: %s", fileType)
	}

	switch config.Type {
	case "local":
		fullPath := filepath.Join(config.Local.Path, path)
		return os.Open(fullPath)
	case "s3":
		sess := session.Must(session.NewSession(&aws.Config{
			Region: aws.String(config.S3.Region),
		}))

		downloader := s3manager.NewDownloader(sess)

		buf := aws.NewWriteAtBuffer([]byte{})
		_, err := downloader.Download(buf, &s3.GetObjectInput{
			Bucket: aws.String(config.S3.Bucket),
			Key:    aws.String(path),
		})

		return ioutil.NopCloser(bytes.NewReader(buf.Bytes())), err
	default:
		return nil, fmt.Errorf("unsupported storage type: %s", config.Type)
	}
}

func DeleteFile(fileType FileType, path string) error {
	config, ok := storageConfigs[fileType]
	if !ok {
		return fmt.Errorf("invalid file type: %s", fileType)
	}

	switch config.Type {
	case "local":
		fullPath := filepath.Join(config.Local.Path, path)
		return os.Remove(fullPath)
	case "s3":
		sess := session.Must(session.NewSession(&aws.Config{
			Region: aws.String(config.S3.Region),
		}))

		deleter := s3.New(sess)

		_, err := deleter.DeleteObject(&s3.DeleteObjectInput{
			Bucket: aws.String(config.S3.Bucket),
			Key:    aws.String(path),
		})

		return err
	default:
		return fmt.Errorf("unsupported storage type: %s", config.Type)
	}
}
