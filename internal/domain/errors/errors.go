package errors

import (
    "errors"
    "fmt"
)

var (
    ErrFileNotFound    = errors.New("file not found")
    ErrPathNotFound    = errors.New("path not found")
    ErrInvalidOperation = errors.New("invalid operation")
)

type FileManagerError struct {
    Op  string
    Path string
    Err error
}

func (e *FileManagerError) Error() string {
    return fmt.Sprintf("operation %s failed for path %s: %v", e.Op, e.Path, e.Err)
}