//go:build !ci

package certs

import "embed"

type FileSystem interface {
	ReadFile(name string) ([]byte, error)
}

//go:embed  "client-cert.pem" "client-key.pem" "ca-cert.pem"
var f embed.FS

var Certs FileSystem = f
