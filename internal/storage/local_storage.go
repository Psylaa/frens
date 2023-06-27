// local_storage.go
package storage

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/bwoff11/frens/internal/config"
	"github.com/google/uuid"
)

type LocalStorage struct {
	Path string
}

func newLocalStorage(cfg *config.StorageDetails) *LocalStorage {
	return &LocalStorage{
		Path: cfg.Local.Path,
	}
}

func (ls *LocalStorage) SaveFile(data []byte, filename string) error {
	fullPath := filepath.Join(ls.Path, filename)
	return ioutil.WriteFile(fullPath, data, 0644)
}

func (ls *LocalStorage) LoadFile(id uuid.UUID) (io.ReadCloser, error) {
	filename, err := ls.GetFilenameByID(id)
	if err != nil {
		return nil, err
	}
	file, err := os.Open(filepath.Join(ls.Path, filename))
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (ls *LocalStorage) DeleteFile(id uuid.UUID) error {
	filename, err := ls.GetFilenameByID(id)
	if err != nil {
		return err
	}
	return os.Remove(filepath.Join(ls.Path, filename))
}

func (ls *LocalStorage) GetFilenameByID(id uuid.UUID) (string, error) {
	matches, err := filepath.Glob(filepath.Join(ls.Path, id.String()+".*"))
	if err != nil {
		return "", err
	}
	if len(matches) == 0 {
		return "", fmt.Errorf("no file found for ID %s", id)
	}
	_, filename := filepath.Split(matches[0]) // Get the filename without the directory path
	return filename, nil
}
