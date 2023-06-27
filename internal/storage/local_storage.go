// local_storage.go
package storage

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/bwoff11/frens/internal/config"
	"github.com/google/uuid"
	"github.com/h2non/filetype"
)

type LocalStorage struct {
	Path string
}

func newLocalStorage(cfg *config.StorageDetails) *LocalStorage {
	return &LocalStorage{
		Path: cfg.Local.Path,
	}
}

func (ls *LocalStorage) DetectFileType(data []byte) config.FileType {
	buf, _ := ioutil.ReadFile(string(data))
	kind, _ := filetype.Match(buf)

	switch kind.MIME.Type {
	case "image":
		return config.Image
	case "video":
		return config.Video
	case "audio":
		return config.Audio
	default:
		return config.Other
	}
}

func (ls *LocalStorage) SaveFile(id uuid.UUID, data []byte) error {
	log.Println("desu")
	log.Println("desu")
	log.Println("desu")
	filePath := ls.filePath(id)
	return ioutil.WriteFile(filePath, data, 0644)
}

func (ls *LocalStorage) LoadFile(id uuid.UUID) (io.ReadCloser, error) {
	filePath := ls.filePath(id)
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (ls *LocalStorage) DeleteFile(id uuid.UUID) error {
	filePath := ls.filePath(id)
	return os.Remove(filePath)
}

// filePath constructs a file path based on the id and file type.
func (ls *LocalStorage) filePath(id uuid.UUID) string {
	dir := filepath.Join(ls.Path, id.String())
	return filepath.Join(dir, id.String())
}
