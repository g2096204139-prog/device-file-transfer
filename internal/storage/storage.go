package storage

import (
	"io"
	"time"
)

type FileInfo struct {
	Filename   string    `json:"filename"`
	Size       int64     `json:"size"`
	UploadedAt time.Time `json:"uploaded_at"`
}

type Storage interface {
	Save(reader io.Reader, filename string) (FileInfo, error)
	List() ([]FileInfo, error)
	Open(filename string) (io.ReadSeekCloser, FileInfo, error)
	Delete(filename string) error
}
