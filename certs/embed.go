//go:build !ci

package certs

import "embed"

//go:embed  "client-cert.pem" "client-key.pem" "ca-cert.pem"
var f embed.FS

var Certs = f
