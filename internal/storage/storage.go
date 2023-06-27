package storage

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/bwoff11/frens/internal/config"
	"github.com/bwoff11/frens/internal/logger"
)

var basePath string

type FileType string
type StorageType string

const (
	FileTypeImage       FileType = "profile_picture"
	FileTypeVideo       FileType = "user_banner"
	FileTypeStatusFile  FileType = "status_media/file"
	FileTypeStatusImage FileType = "status_media/image"
	FileTypeStatusVideo FileType = "status_media/video"
)

const (
	StorageTypeLocal StorageType = "local"
	StorageTypeS3    StorageType = "s3"
)

func InitStorage(cfg *config.Config) error {
	logger.Log.Info().Msg("Initializing storage")

	// Parse the config and set the base path
	basePath = cfg.Storage.Local.Path
	logger.Log.Info().Str("path", basePath).Msg("Set base path for storage to:" + basePath)

	// Check if the base path exists
	_, err := os.Stat(basePath)
	if os.IsNotExist(err) {
		return fmt.Errorf("base path %s does not exist", basePath)
	} else if err != nil {
		return fmt.Errorf("error checking base path %s: %v", basePath, err)
	}

	// Create folder for each file type if it doesn't exist
	for _, fileType := range []FileType{FileTypeImage, FileTypeVideo, FileTypeStatusFile, FileTypeStatusImage, FileTypeStatusVideo} {
		err := os.MkdirAll(filepath.Join(basePath, string(fileType)), 0755)
		if err != nil {
			return err
		}
	}

	// Test creating and deleting a file
	err = ioutil.WriteFile(filepath.Join(basePath, "test.txt"), []byte("test"), 0644)
	if err != nil {
		return err
	} else {
		logger.Log.Info().Msg("Created test file")
	}
	err = os.Remove(filepath.Join(basePath, "test.txt"))
	if err != nil {
		return err
	} else {
		logger.Log.Info().Msg("Deleted test file")
	}

	return nil
}

func SaveFile(path string, data []byte) error {
	fullPath := filepath.Join(basePath, path)
	return ioutil.WriteFile(fullPath, data, 0644)
}

func LoadFile(path string) ([]byte, error) {
	fullPath := filepath.Join(basePath, path)
	return ioutil.ReadFile(fullPath)
}

func DeleteFile(path string) error {
	fullPath := filepath.Join(basePath, path)
	return os.Remove(fullPath)
}
