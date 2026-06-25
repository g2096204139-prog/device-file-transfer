package util

import (
	"errors"
	"path/filepath"
	"regexp"
	"strings"
)

var safeNamePattern = regexp.MustCompile(`[^a-zA-Z0-9._\-\p{Han}]`)

var blockedExtensions = map[string]bool{
	".exe": true,
	".bat": true,
	".cmd": true,
	".com": true,
	".scr": true,
	".ps1": true,
	".sh":  true,
	".msi": true,
	".dll": true,
	".vbs": true,
	".js":  true,
	".jar": true,
	".apk": true,
	".app": true,
}

func SanitizeFilename(input string) (string, error) {
	trimmed := strings.TrimSpace(input)
	if strings.Contains(trimmed, "..") || strings.ContainsAny(trimmed, `/\`) {
		return "", errors.New("filename contains unsafe sequence")
	}

	name := filepath.Base(trimmed)
	name = safeNamePattern.ReplaceAllString(name, "_")

	if name == "" || name == "." || name == ".." {
		return "", errors.New("invalid filename")
	}

	if IsBlockedExtension(name) {
		return "", errors.New("blocked file extension")
	}

	return name, nil
}

func IsBlockedExtension(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return blockedExtensions[ext]
}
