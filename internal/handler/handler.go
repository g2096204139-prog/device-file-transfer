package handler

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/g2096204139-prog/device-file-transfer/internal/auth"
	"github.com/g2096204139-prog/device-file-transfer/internal/config"
	"github.com/g2096204139-prog/device-file-transfer/internal/logger"
	"github.com/g2096204139-prog/device-file-transfer/internal/service"
)

type Handler struct {
	service *service.FileService
	cfg     config.Config
	logger  *logger.AppLogger
}

func NewHandler(fileService *service.FileService, cfg config.Config, appLogger *logger.AppLogger) *Handler {
	return &Handler{
		service: fileService,
		cfg:     cfg,
		logger:  appLogger,
	}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.index)
	mux.HandleFunc("/api/health", h.health)
	mux.HandleFunc("/api/server-info", h.serverInfo)
	mux.HandleFunc("/api/files", h.files)
	mux.HandleFunc("/api/upload", h.upload)
	mux.HandleFunc("/api/download/", h.download)

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
}

func (h *Handler) requireAuth(w http.ResponseWriter, r *http.Request) bool {
	if auth.VerifyToken(r, h.cfg.AccessToken, h.cfg.AuthEnabled) {
		return true
	}

	h.logger.Warn("auth_failed", r.RemoteAddr+" "+r.URL.Path)
	writeError(w, http.StatusUnauthorized, "unauthorized", "UNAUTHORIZED")
	return false
}

func (h *Handler) index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		writeError(w, http.StatusNotFound, "not found", "NOT_FOUND")
		return
	}

	content, err := os.ReadFile("web/templates/index.html")
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to load page", "PAGE_LOAD_FAILED")
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write(content)
}

func (h *Handler) health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, APIResponse{
		Success: true,
		Message: "ok",
		Data: map[string]string{
			"status": "healthy",
		},
	})
}

func (h *Handler) serverInfo(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, APIResponse{
		Success: true,
		Message: "server info",
		Data: map[string]interface{}{
			"name":               "Device File Transfer",
			"version":            "1.0.6",
			"host":               h.cfg.ServerHost,
			"port":               h.cfg.ServerPort,
			"lan_ips":            localIPv4s(),
			"auth_enabled":       h.cfg.AuthEnabled,
			"max_upload_size_mb": h.cfg.MaxUploadSizeMB,
			"allow_delete":       h.cfg.AllowDelete,
		},
	})
}

func (h *Handler) files(w http.ResponseWriter, r *http.Request) {
	if !h.requireAuth(w, r) {
		return
	}

	switch r.Method {
	case http.MethodGet:
		files, err := h.service.ListFiles()
		if err != nil {
			writeError(w, http.StatusInternalServerError, "failed to list files", "LIST_FAILED")
			return
		}
		writeJSON(w, http.StatusOK, APIResponse{Success: true, Message: "files", Data: files})
	case http.MethodDelete:
		filename := r.URL.Query().Get("filename")
		if filename == "" {
			writeError(w, http.StatusBadRequest, "filename is required", "FILENAME_REQUIRED")
			return
		}
		h.deleteByFilename(w, filename)
	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed", "METHOD_NOT_ALLOWED")
	}
}

func (h *Handler) upload(w http.ResponseWriter, r *http.Request) {
	if !h.requireAuth(w, r) {
		return
	}

	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed", "METHOD_NOT_ALLOWED")
		return
	}

	maxBytes := h.cfg.MaxUploadBytes()
	if maxBytes > 0 {
		r.Body = http.MaxBytesReader(w, r.Body, maxBytes+10*1024*1024)
	}

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		writeError(w, http.StatusBadRequest, "invalid multipart form or file too large", "INVALID_UPLOAD")
		return
	}

	fileHeaders := r.MultipartForm.File["files"]
	if len(fileHeaders) == 0 {
		if single := r.MultipartForm.File["file"]; len(single) > 0 {
			fileHeaders = single
		}
	}

	if len(fileHeaders) == 0 {
		writeError(w, http.StatusBadRequest, "no files uploaded", "NO_FILES")
		return
	}

	results := make([]service.UploadResult, 0, len(fileHeaders))
	for _, fileHeader := range fileHeaders {
		results = append(results, h.service.SaveMultipartFile(fileHeader))
	}

	writeJSON(w, http.StatusOK, APIResponse{
		Success: true,
		Message: "upload processed",
		Data: map[string]interface{}{
			"results": results,
		},
	})
}

func (h *Handler) download(w http.ResponseWriter, r *http.Request) {
	if !h.requireAuth(w, r) {
		return
	}

	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed", "METHOD_NOT_ALLOWED")
		return
	}

	filename := strings.TrimPrefix(r.URL.Path, "/api/download/")
	filename = path.Clean("/" + filename)
	filename = strings.TrimPrefix(filename, "/")

	file, info, err := h.service.OpenFile(filename)
	if err != nil {
		writeError(w, http.StatusNotFound, "file not found", "FILE_NOT_FOUND")
		return
	}
	defer file.Close()

	w.Header().Set("Accept-Ranges", "bytes")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%q", info.Filename))
	http.ServeContent(w, r, info.Filename, info.UploadedAt, file)
}

func (h *Handler) deleteByFilename(w http.ResponseWriter, filename string) {
	if err := h.service.DeleteFile(filename); err != nil {
		writeError(w, http.StatusBadRequest, err.Error(), "DELETE_FAILED")
		return
	}

	writeJSON(w, http.StatusOK, APIResponse{
		Success: true,
		Message: "delete success",
		Data: map[string]string{
			"filename": filename,
		},
	})
}

func localIPv4s() []string {
	var result []string

	interfaces, err := net.Interfaces()
	if err != nil {
		return result
	}

	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			var ip net.IP

			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip == nil || ip.To4() == nil {
				continue
			}

			result = append(result, ip.String())
		}
	}

	return result
}

func decodeJSON(r *http.Request, target interface{}) error {
	return json.NewDecoder(r.Body).Decode(target)
}
