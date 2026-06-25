package util

import "testing"

func TestSanitizeFilenameValid(t *testing.T) {
	name, err := SanitizeFilename("照片-01.jpg")
	if err != nil {
		t.Fatalf("expected valid filename, got error: %v", err)
	}
	if name != "照片-01.jpg" {
		t.Fatalf("unexpected sanitized name: %s", name)
	}
}

func TestSanitizeFilenameRejectsTraversal(t *testing.T) {
	_, err := SanitizeFilename("../secret.txt")
	if err == nil {
		t.Fatal("expected traversal-like filename to be rejected")
	}
}

func TestSanitizeFilenameRejectsBlockedExtension(t *testing.T) {
	_, err := SanitizeFilename("tool.exe")
	if err == nil {
		t.Fatal("expected .exe to be blocked")
	}
}

func TestSanitizeFilenameReplacesUnsafeChars(t *testing.T) {
	name, err := SanitizeFilename("my file!.txt")
	if err != nil {
		t.Fatalf("expected sanitized filename, got error: %v", err)
	}
	if name != "my_file_.txt" {
		t.Fatalf("unexpected sanitized name: %s", name)
	}
}
