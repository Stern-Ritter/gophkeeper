package client

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"

	embed "github.com/Stern-Ritter/gophkeeper/certs"

	"google.golang.org/grpc/credentials"
)

func GetTransportCredentials() (credentials.TransportCredentials, error) {
	certPEM, err := embed.Certs.ReadFile("client-cert.pem")
	if err != nil {
		return nil, fmt.Errorf("error reading client-cert.pem: %v", err)
	}

	keyPEM, err := embed.Certs.ReadFile("client-key.pem")
	if err != nil {
		return nil, fmt.Errorf("error reading client-key.pem: %v", err)
	}

	cert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		return nil, fmt.Errorf("error loading client cert: %v", err)
	}

	caPEM, err := embed.Certs.ReadFile("ca-cert.pem")
	if err != nil {
		return nil, fmt.Errorf("error reading ca-cert.pem: %v", err)
	}

	caCertPool := x509.NewCertPool()
	if !caCertPool.AppendCertsFromPEM(caPEM) {
		return nil, fmt.Errorf("failed to append ca cert to pool")
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}

	return credentials.NewTLS(tlsConfig), nil
}
