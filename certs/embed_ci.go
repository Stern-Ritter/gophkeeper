//go:build ci

package certs

import (
	"io/fs"
	"os"
	"path/filepath"
)

type mockFS struct {
	dir string
}

func NewMockFS(dir string) *mockFS {
	return &mockFS{dir: dir}
}

func (r *mockFS) Open(name string) (fs.File, error) {
	fullPath := filepath.Join(r.dir, name)
	file, err := os.Open(fullPath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (r *mockFS) ReadFile(name string) ([]byte, error) {
	fullPath := filepath.Join(r.dir, name)
	data, err := os.ReadFile(fullPath)
	if err != nil {
		return nil, err
	}
	return data, nil
}

var Certs fs.ReadFileFS = NewMockFS("../../../testdata/certs")
