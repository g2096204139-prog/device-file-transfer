package service

import (
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/g2096204139-prog/device-file-transfer/internal/config"
	"github.com/g2096204139-prog/device-file-transfer/internal/logger"
	"github.com/g2096204139-prog/device-file-transfer/internal/storage"
	"github.com/g2096204139-prog/device-file-transfer/internal/util"
)

type FileService struct {
	store  storage.Storage
	cfg    config.Config
	logger *logger.AppLogger
}

type UploadResult struct {
	Filename string `json:"filename"`
	Size     int64  `json:"size"`
	Status   string `json:"status"`
	Error    string `json:"error,omitempty"`
}

type FileDTO struct {
	Filename   string `json:"filename"`
	Size       int64  `json:"size"`
	UploadedAt string `json:"uploaded_at"`
}

func NewFileService(store storage.Storage, cfg config.Config, appLogger *logger.AppLogger) *FileService {
	return &FileService{store: store, cfg: cfg, logger: appLogger}
}

func (s *FileService) ListFiles() ([]FileDTO, error) {
	files, err := s.store.List()
	if err != nil {
		return nil, err
	}

	result := make([]FileDTO, 0, len(files))
	for _, file := range files {
		result = append(result, toDTO(file))
	}
	return result, nil
}

func (s *FileService) SaveMultipartFile(fileHeader *multipart.FileHeader) UploadResult {
	safeName, err := util.SanitizeFilename(fileHeader.Filename)
	if err != nil {
		s.logger.Warn("upload_rejected", fileHeader.Filename+": "+err.Error())
		return UploadResult{Filename: fileHeader.Filename, Status: "failed", Error: err.Error()}
	}

	maxBytes := s.cfg.MaxUploadBytes()
	if maxBytes > 0 && fileHeader.Size > maxBytes {
		s.logger.Warn("upload_rejected", safeName+": file too large")
		return UploadResult{Filename: safeName, Status: "failed", Error: "file too large"}
	}

	file, err := fileHeader.Open()
	if err != nil {
		s.logger.Error("upload_open_failed", err)
		return UploadResult{Filename: safeName, Status: "failed", Error: "failed to open upload"}
	}
	defer file.Close()

	var reader io.Reader = file
	if maxBytes > 0 {
		reader = http.MaxBytesReader(nil, file, maxBytes)
	}

	info, err := s.store.Save(reader, safeName)
	if err != nil {
		s.logger.Error("upload_save_failed", err)
		return UploadResult{Filename: safeName, Status: "failed", Error: "failed to save file"}
	}

	s.logger.Info("upload_success", info.Filename)
	return UploadResult{Filename: info.Filename, Size: info.Size, Status: "success"}
}

func (s *FileService) OpenFile(filename string) (io.ReadSeekCloser, storage.FileInfo, error) {
	safeName, err := util.SanitizeFilename(filename)
	if err != nil {
		return nil, storage.FileInfo{}, err
	}

	file, info, err := s.store.Open(safeName)
	if err != nil {
		return nil, storage.FileInfo{}, err
	}

	s.logger.Info("download", safeName)
	return file, info, nil
}

func (s *FileService) DeleteFile(filename string) error {
	if !s.cfg.AllowDelete {
		return errors.New("delete disabled")
	}

	safeName, err := util.SanitizeFilename(filename)
	if err != nil {
		return err
	}

	if err := s.store.Delete(safeName); err != nil {
		return err
	}

	s.logger.Info("delete", safeName)
	return nil
}

func toDTO(info storage.FileInfo) FileDTO {
	return FileDTO{
		Filename:   info.Filename,
		Size:       info.Size,
		UploadedAt: info.UploadedAt.Format(time.RFC3339),
	}
}
