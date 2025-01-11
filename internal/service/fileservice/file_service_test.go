package fileservice

import (
    "os"
    "path/filepath"
    "testing"
)

type mockLogger struct{}

func (m *mockLogger) Info(msg string, keysAndValues ...interface{})  {}
func (m *mockLogger) Error(msg string, keysAndValues ...interface{}) {}

func TestFileService(t *testing.T) {
    // Create temporary directory for testing
    tempDir, err := os.MkdirTemp("", "filemanager-test-")
    if err != nil {
        t.Fatalf("Failed to create temp directory: %v", err)
    }
    defer os.RemoveAll(tempDir)

    service := New(tempDir, &mockLogger{})

    // Test file creation
    t.Run("Create", func(t *testing.T) {
        err := service.Create("test.txt")
        if err != nil {
            t.Errorf("Failed to create file: %v", err)
        }

        if _, err := os.Stat(filepath.Join(tempDir, "test.txt")); os.IsNotExist(err) {
            t.Error("File was not created")
        }
    })

    // Add more tests for List, Delete, Copy, and Move operations
}