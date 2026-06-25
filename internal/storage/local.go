package storage

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type LocalStorage struct {
	dir string
}

func NewLocalStorage(dir string) (*LocalStorage, error) {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}
	return &LocalStorage{dir: dir}, nil
}

func (s *LocalStorage) Save(reader io.Reader, filename string) (FileInfo, error) {
	targetPath := s.uniquePath(filename)

	file, err := os.OpenFile(targetPath, os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0644)
	if err != nil {
		return FileInfo{}, err
	}
	defer file.Close()

	if _, err := io.Copy(file, reader); err != nil {
		_ = os.Remove(targetPath)
		return FileInfo{}, err
	}

	stat, err := file.Stat()
	if err != nil {
		return FileInfo{}, err
	}

	return FileInfo{
		Filename:   filepath.Base(targetPath),
		Size:       stat.Size(),
		UploadedAt: stat.ModTime(),
	}, nil
}

func (s *LocalStorage) List() ([]FileInfo, error) {
	entries, err := os.ReadDir(s.dir)
	if err != nil {
		return nil, err
	}

	files := make([]FileInfo, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() || isHiddenSystemFile(entry.Name()) {
			continue
		}

		info, err := entry.Info()
		if err != nil {
			continue
		}

		files = append(files, FileInfo{
			Filename:   entry.Name(),
			Size:       info.Size(),
			UploadedAt: info.ModTime(),
		})
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].UploadedAt.After(files[j].UploadedAt)
	})

	return files, nil
}

func (s *LocalStorage) Open(filename string) (io.ReadSeekCloser, FileInfo, error) {
	if isHiddenSystemFile(filename) {
		return nil, FileInfo{}, errors.New("file not found")
	}

	path := filepath.Join(s.dir, filename)

	file, err := os.Open(path)
	if err != nil {
		return nil, FileInfo{}, err
	}

	stat, err := file.Stat()
	if err != nil {
		_ = file.Close()
		return nil, FileInfo{}, err
	}

	return file, FileInfo{
		Filename:   filename,
		Size:       stat.Size(),
		UploadedAt: stat.ModTime(),
	}, nil
}

func (s *LocalStorage) Delete(filename string) error {
	if isHiddenSystemFile(filename) {
		return errors.New("cannot delete system file")
	}

	path := filepath.Join(s.dir, filename)
	return os.Remove(path)
}

func (s *LocalStorage) uniquePath(filename string) string {
	path := filepath.Join(s.dir, filename)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return path
	}

	ext := filepath.Ext(filename)
	stem := strings.TrimSuffix(filename, ext)

	for i := 1; ; i++ {
		candidate := filepath.Join(s.dir, fmt.Sprintf("%s_%d%s", stem, i, ext))
		if _, err := os.Stat(candidate); os.IsNotExist(err) {
			return candidate
		}
	}
}

func isHiddenSystemFile(filename string) bool {
	switch strings.ToLower(filename) {
	case ".gitkeep", ".ds_store", "thumbs.db":
		return true
	default:
		return false
	}
}

func FormatTime(t time.Time) string {
	return t.Format(time.RFC3339)
}
