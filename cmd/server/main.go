package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/g2096204139-prog/device-file-transfer/internal/config"
	"github.com/g2096204139-prog/device-file-transfer/internal/handler"
	"github.com/g2096204139-prog/device-file-transfer/internal/logger"
	"github.com/g2096204139-prog/device-file-transfer/internal/service"
	"github.com/g2096204139-prog/device-file-transfer/internal/storage"
)

func main() {
	cfg := config.Load()

	appLogger, err := logger.New(cfg.LogDir)
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
	defer appLogger.Close()

	store, err := storage.NewLocalStorage(cfg.UploadDir)
	if err != nil {
		log.Fatalf("failed to initialize storage: %v", err)
	}

	fileService := service.NewFileService(store, cfg, appLogger)
	appHandler := handler.NewHandler(fileService, cfg, appLogger)

	mux := http.NewServeMux()
	appHandler.RegisterRoutes(mux)

	addr := fmt.Sprintf("%s:%s", cfg.ServerHost, cfg.ServerPort)
	log.Printf("Device File Transfer server running at http://%s", addr)
	log.Printf("Access token is configured. Change ACCESS_TOKEN before real use.")

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
