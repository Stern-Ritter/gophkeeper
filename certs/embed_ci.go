//go:build ci

package certs

import (
	"errors"
	"io/fs"
)

type mockFS struct{}

func (mockFS) Open(name string) (fs.File, error) {
	return nil, errors.New("mockFS does not contain files")
}

func (mockFS) ReadFile(name string) ([]byte, error) {
	return nil, errors.New("mockFS does not contain files")
}

var Certs fs.ReadFileFS = mockFS{}
