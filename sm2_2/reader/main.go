package main

import (
	"fmt"
	"io"
	"os"
)

// Task 2
func main() {
	r := NewReader()

	r.Read("file.go")
}

type Reader interface {
	Read(path string) error
	Display() error
}

func NewReader() Reader {
	return &reader{}
}

type reader struct {
	r io.Reader
}

func (r *reader) Read(path string) error {
	stat, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("file not found")
	}

	if stat.IsDir() {
		return fmt.Errorf("it is not a file: %w", err)
	}

	b, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	fmt.Printf("size: %d modified: %v\n\n", stat.Size(), stat.ModTime())
	fmt.Println(string(b))

	return nil
}

func (r *reader) Display() error {
	return nil
}
