package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

type AppLogger struct {
	file *os.File
	log  *log.Logger
}

func New(logDir string) (*AppLogger, error) {
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return nil, err
	}

	filename := time.Now().Format("2006-01-02") + ".log"
	path := filepath.Join(logDir, filename)

	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	return &AppLogger{
		file: file,
		log:  log.New(file, "", log.LstdFlags),
	}, nil
}

func (l *AppLogger) Info(action, detail string) {
	if l == nil || l.log == nil {
		return
	}
	l.log.Printf("INFO action=%s detail=%q", action, detail)
}

func (l *AppLogger) Warn(action, detail string) {
	if l == nil || l.log == nil {
		return
	}
	l.log.Printf("WARN action=%s detail=%q", action, detail)
}

func (l *AppLogger) Error(action string, err error) {
	if l == nil || l.log == nil {
		return
	}
	l.log.Printf("ERROR action=%s error=%q", action, fmt.Sprint(err))
}

func (l *AppLogger) Close() error {
	if l == nil || l.file == nil {
		return nil
	}
	return l.file.Close()
}
