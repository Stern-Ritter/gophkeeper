package client

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Stern-Ritter/gophkeeper/certs"
)

type mockFS struct {
	files map[string][]byte
}

func (m mockFS) ReadFile(name string) ([]byte, error) {
	if data, ok := m.files[name]; ok {
		return data, nil
	}
	return nil, fmt.Errorf("file not found: %s", name)
}

func mockCertsFS(files map[string][]byte) {
	certs.Certs = mockFS{files: files}
}

func readFileContent(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("error reading file %s: %w", filePath, err)
	}
	return string(content), nil
}

func TestGetTransportCredentials_Success(t *testing.T) {
	clientCert, err := readFileContent("../../../testdata/certs/client-cert.pem")
	require.NoError(t, err, "unexpected error reading client cert")
	clientKey, err := readFileContent("../../../testdata/certs/client-key.pem")
	require.NoError(t, err, "unexpected error reading client key")
	caCert, err := readFileContent("../../../testdata/certs/ca-cert.pem")
	require.NoError(t, err, "unexpected error reading ca cert")

	mockCertsFS(map[string][]byte{
		"client-cert.pem": []byte(clientCert),
		"client-key.pem":  []byte(clientKey),
		"ca-cert.pem":     []byte(caCert),
	})

	creds, err := GetTransportCredentials()
	require.NoError(t, err, "unexpected error getting transport credentials")
	require.NotNil(t, creds, "should return transport credentials")
}

func TestGetTransportCredentials_ClientCertMissing(t *testing.T) {
	clientKey, err := readFileContent("../../../testdata/certs/client-key.pem")
	require.NoError(t, err, "unexpected error reading client key")
	caCert, err := readFileContent("../../../testdata/certs/ca-cert.pem")
	require.NoError(t, err, "unexpected error reading ca cert")

	mockCertsFS(map[string][]byte{
		"client-key.pem": []byte(clientKey),
		"ca-cert.pem":    []byte(caCert),
	})

	_, err = GetTransportCredentials()
	require.Error(t, err, "expected error missing client cert")
	assert.Contains(t, err.Error(), "error reading client-cert.pem")
}

func TestGetTransportCredentials_ClientKeyMissing(t *testing.T) {
	clientCert, err := readFileContent("../../../testdata/certs/client-cert.pem")
	require.NoError(t, err, "unexpected error reading client cert")
	caCert, err := readFileContent("../../../testdata/certs/ca-cert.pem")
	require.NoError(t, err, "unexpected error reading ca cert")

	mockCertsFS(map[string][]byte{
		"client-cert.pem": []byte(clientCert),
		"ca-cert.pem":     []byte(caCert),
	})

	_, err = GetTransportCredentials()
	require.Error(t, err, "expected error missing client key")
	assert.Contains(t, err.Error(), "error reading client-key.pem")
}

func TestGetTransportCredentials_CACertMissing(t *testing.T) {
	clientCert, err := readFileContent("../../../testdata/certs/client-cert.pem")
	require.NoError(t, err, "unexpected error reading client cert")
	clientKey, err := readFileContent("../../../testdata/certs/client-key.pem")
	require.NoError(t, err, "unexpected error reading client key")

	mockCertsFS(map[string][]byte{
		"client-cert.pem": []byte(clientCert),
		"client-key.pem":  []byte(clientKey),
	})

	_, err = GetTransportCredentials()
	require.Error(t, err, "expected error missing ca cert")
	assert.Contains(t, err.Error(), "error reading ca-cert.pem")
}

func TestGetTransportCredentials_InvalidClientCert(t *testing.T) {
	clientKey, err := readFileContent("../../../testdata/certs/client-key.pem")
	require.NoError(t, err, "unexpected error reading client key")
	caCert, err := readFileContent("../../../testdata/certs/ca-cert.pem")
	require.NoError(t, err, "unexpected error reading ca cert")

	mockCertsFS(map[string][]byte{
		"client-cert.pem": []byte("invalid cert"),
		"client-key.pem":  []byte(clientKey),
		"ca-cert.pem":     []byte(caCert),
	})

	_, err = GetTransportCredentials()
	require.Error(t, err, "expected error loading client cert")
	assert.Contains(t, err.Error(), "error loading client cert")
}

func TestGetTransportCredentials_InvalidClientKey(t *testing.T) {
	clientCert, err := readFileContent("../../../testdata/certs/client-cert.pem")
	require.NoError(t, err, "unexpected error reading client cert")
	caCert, err := readFileContent("../../../testdata/certs/ca-cert.pem")
	require.NoError(t, err, "unexpected error reading ca cert")

	mockCertsFS(map[string][]byte{
		"client-cert.pem": []byte(clientCert),
		"client-key.pem":  []byte("invalid key"),
		"ca-cert.pem":     []byte(caCert),
	})

	_, err = GetTransportCredentials()
	require.Error(t, err, "expected error loading client key")
	assert.Contains(t, err.Error(), "error loading client cert: tls: failed to find any PEM data in key input")
}

func TestGetTransportCredentials_InvalidCACert(t *testing.T) {
	clientCert, err := readFileContent("../../../testdata/certs/client-cert.pem")
	require.NoError(t, err, "unexpected error reading client cert")
	clientKey, err := readFileContent("../../../testdata/certs/client-key.pem")
	require.NoError(t, err, "unexpected error reading client key")

	mockCertsFS(map[string][]byte{
		"client-cert.pem": []byte(clientCert),
		"client-key.pem":  []byte(clientKey),
		"ca-cert.pem":     []byte("invalid cert"),
	})

	_, err = GetTransportCredentials()
	require.Error(t, err, "expected error loading ca cert")
	assert.Contains(t, err.Error(), "failed to append ca cert to pool")
}
