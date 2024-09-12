//go:build !ci

package certs

import (
	"embed"
	"io/fs"
)

type FileSystem interface {
	Open(name string) (fs.File, error)
	ReadDir(name string) ([]fs.DirEntry, error)
	ReadFile(name string) ([]byte, error)
}

//go:embed  "client-cert.pem" "client-key.pem" "ca-cert.pem"
var f embed.FS

var Certs FileSystem = f
