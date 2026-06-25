package storage

import (
	"strings"
	"testing"
)

func TestLocalStorageSaveListOpenDelete(t *testing.T) {
	dir := t.TempDir()
	store, err := NewLocalStorage(dir)
	if err != nil {
		t.Fatalf("failed to create storage: %v", err)
	}

	info, err := store.Save(strings.NewReader("hello"), "hello.txt")
	if err != nil {
		t.Fatalf("save failed: %v", err)
	}

	if info.Filename != "hello.txt" {
		t.Fatalf("unexpected filename: %s", info.Filename)
	}

	files, err := store.List()
	if err != nil {
		t.Fatalf("list failed: %v", err)
	}

	if len(files) != 1 {
		t.Fatalf("expected 1 file, got %d", len(files))
	}

	file, _, err := store.Open("hello.txt")
	if err != nil {
		t.Fatalf("open failed: %v", err)
	}
	_ = file.Close()

	if err := store.Delete("hello.txt"); err != nil {
		t.Fatalf("delete failed: %v", err)
	}
}

func TestLocalStorageDuplicateName(t *testing.T) {
	dir := t.TempDir()
	store, err := NewLocalStorage(dir)
	if err != nil {
		t.Fatalf("failed to create storage: %v", err)
	}

	first, err := store.Save(strings.NewReader("one"), "file.txt")
	if err != nil {
		t.Fatalf("first save failed: %v", err)
	}

	second, err := store.Save(strings.NewReader("two"), "file.txt")
	if err != nil {
		t.Fatalf("second save failed: %v", err)
	}

	if first.Filename == second.Filename {
		t.Fatal("expected duplicate filename to be renamed")
	}
}

func TestLocalStorageIgnoresSystemFiles(t *testing.T) {
	dir := t.TempDir()
	store, err := NewLocalStorage(dir)
	if err != nil {
		t.Fatalf("failed to create storage: %v", err)
	}

	if _, err := store.Save(strings.NewReader("placeholder"), ".gitkeep"); err != nil {
		t.Fatalf("save system file placeholder failed: %v", err)
	}
	if _, err := store.Save(strings.NewReader("hello"), "hello.txt"); err != nil {
		t.Fatalf("save normal file failed: %v", err)
	}

	files, err := store.List()
	if err != nil {
		t.Fatalf("list failed: %v", err)
	}

	if len(files) != 1 || files[0].Filename != "hello.txt" {
		t.Fatalf("expected only hello.txt in list, got %+v", files)
	}

	if _, _, err := store.Open(".gitkeep"); err == nil {
		t.Fatal("expected .gitkeep open to be rejected")
	}

	if err := store.Delete(".gitkeep"); err == nil {
		t.Fatal("expected .gitkeep delete to be rejected")
	}
}
