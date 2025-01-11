// file.go - This file defines data structures related to file information and operations.

package models // This declares that the code belongs to the 'models' package.

import (
	"os" // Importing the 'os' package to use file-related types.
	"time" // Importing the 'time' package to use time-related types.
)

// FileInfo struct holds information about a file.
type FileInfo struct {
	Name     string    // The name of the file (e.g., "document.txt").
	Size     int64     // The size of the file in bytes.
	Mode     os.FileMode // The file's permissions and mode (e.g., read/write).
	ModifiedAt  time.Time // The last time the file's mode was changed.
	IsDir    bool      // A boolean indicating if the file is a directory (true) or not (false).
	
}

// FileOperation struct represents an operation that can be performed on a file.
type FileOperation struct {
	Type       string // The type of operation (e.g., "create", "delete", "move").
	Path       string // The current path of the file.
	TargetPath string // The target path where the file should be moved or copied.
}
