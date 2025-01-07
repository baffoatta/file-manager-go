package fileservice

import (
    "io"
    "os"
    "path/filepath"

    "github.com/baffoatta/filemanager/internal/domain/errors"
    "github.com/baffoatta/filemanager/internal/domain/models"
)

// FileService is a struct that holds the base directory and a logger for logging messages.
type FileService struct {
    baseDir string // The base directory where files are managed
    logger  Logger  // Logger interface for logging information and errors
}

// Logger is an interface that defines methods for logging messages.
type Logger interface {
    Info(msg string, keysAndValues ...interface{}) // Method to log informational messages
    Error(msg string, keysAndValues ...interface{}) // Method to log error messages
}

// New creates a new instance of FileService with the specified base directory and logger.
func New(baseDir string, logger Logger) *FileService {
    return &FileService{
        baseDir: baseDir, // Set the base directory
        logger:  logger,   // Set the logger
    }
}

// List retrieves a list of files in the specified path.
func (s *FileService) List(path string) ([]models.FileInfo, error) {
    fullPath := filepath.Join(s.baseDir, path) // Combine baseDir and path to get the full path
    
    // Read the directory entries
    entries, err := os.ReadDir(fullPath)
    if err != nil {
        // Return an error if reading the directory fails
        return nil, &errors.FileManagerError{
            Op:   "list", // Operation name
            Path: path,   // Path that caused the error
            Err:  err,    // Original error
        }
    }

    var files []models.FileInfo // Slice to hold file information
    for _, entry := range entries {
        info, err := entry.Info() // Get information about the entry
        if err != nil {
            // Log an error if getting file info fails and continue to the next entry
            s.logger.Error("Failed to get file info", "file", entry.Name(), "error", err)
            continue
        }

        // Append the file information to the files slice
        files = append(files, models.FileInfo{
            Name:    info.Name(),    // File name
            Size:    info.Size(),    // File size
            Mode:    info.Mode(),    // File mode (permissions)
            ModTime: info.ModTime(), // Last modification time
            IsDir:   info.IsDir(),   // Whether the entry is a directory
        })
    }

    return files, nil // Return the list of files
}

// Create creates a new file at the specified path.
func (s *FileService) Create(path string) error {
    fullPath := filepath.Join(s.baseDir, path) // Get the full path for the new file
    
    dir := filepath.Dir(fullPath) // Get the directory of the new file
    // Create the directory if it doesn't exist
    if err := os.MkdirAll(dir, 0755); err != nil {
        return &errors.FileManagerError{
            Op:   "create", // Operation name
            Path: path,     // Path that caused the error
            Err:  err,      // Original error
        }
    }

    // Create the new file
    file, err := os.Create(fullPath)
    if err != nil {
        return &errors.FileManagerError{
            Op:   "create", // Operation name
            Path: path,     // Path that caused the error
            Err:  err,      // Original error
        }
    }
    defer file.Close() // Ensure the file is closed when done

    return nil // Return nil if successful
}

// Delete removes the file at the specified path.
func (s *FileService) Delete(path string) error {
    fullPath := filepath.Join(s.baseDir, path) // Get the full path of the file to delete
    
    // Remove the file
    err := os.Remove(fullPath)
    if err != nil {
        return &errors.FileManagerError{
            Op:   "delete", // Operation name
            Path: path,     // Path that caused the error
            Err:  err,      // Original error
        }
    }

    return nil // Return nil if successful
}

// Copy duplicates a file from the source path to the destination path.
func (s *FileService) Copy(src, dst string) error {
    srcPath := filepath.Join(s.baseDir, src) // Get the full source path
    dstPath := filepath.Join(s.baseDir, dst) // Get the full destination path

    // Open the source file for reading
    sourceFile, err := os.Open(srcPath)
    if err != nil {
        return &errors.FileManagerError{
            Op:   "copy", // Operation name
            Path: src,    // Path that caused the error
            Err:  err,    // Original error
        }
    }
    defer sourceFile.Close() // Ensure the source file is closed when done

    // Create destination directory if it doesn't exist
    if err := os.MkdirAll(filepath.Dir(dstPath), 0755); err != nil {
        return &errors.FileManagerError{
            Op:   "copy", // Operation name
            Path: dst,    // Path that caused the error
            Err:  err,    // Original error
        }
    }

    // Create the destination file
    destFile, err := os.Create(dstPath)
    if err != nil {
        return &errors.FileManagerError{
            Op:   "copy", // Operation name
            Path: dst,    // Path that caused the error
            Err:  err,    // Original error
        }
    }
    defer destFile.Close() // Ensure the destination file is closed when done

    // Copy the contents from the source file to the destination file
    if _, err := io.Copy(destFile, sourceFile); err != nil {
        return &errors.FileManagerError{
            Op:   "copy", // Operation name
            Path: src,    // Path that caused the error
            Err:  err,    // Original error
        }
    }

    return nil // Return nil if successful
}

// Move moves a file from the source path to the destination path.
func (s *FileService) Move(src, dst string) error {
    // First, copy the file to the new location
    if err := s.Copy(src, dst); err != nil {
        return err // Return the error if copying fails
    }
    // Then, delete the original file
    return s.Delete(src) // Return the result of the delete operation
}
