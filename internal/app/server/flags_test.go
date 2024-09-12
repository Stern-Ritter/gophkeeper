package server

import (
	"flag"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	config "github.com/Stern-Ritter/gophkeeper/internal/config/server"
)

func TestGetConfig_DefaultConfig(t *testing.T) {
	resetFlags()
	os.Args = []string{
		"cmd",
	}

	defaultCfg := config.ServerConfig{
		URL:               "localhost:8080",
		DatabaseDSN:       "default-dsn",
		FileStoragePath:   "/default-storage-path",
		AuthenticationKey: "default-auth-key",
		EncryptionKey:     "default-encrypt-key",
		TLSCertPath:       "/default-tls-cert",
		TLSKeyPath:        "/default-tls-key",
		ShutdownTimeout:   30,
		LoggerLvl:         "info",
	}

	cfg, err := GetConfig(defaultCfg)
	require.NoError(t, err, "unexpected error getting config")

	assert.Equal(t, "localhost:8080", cfg.URL)
	assert.Equal(t, "default-dsn", cfg.DatabaseDSN)
	assert.Equal(t, "/default-storage-path", cfg.FileStoragePath)
	assert.Equal(t, "default-auth-key", cfg.AuthenticationKey)
	assert.Equal(t, "default-encrypt-key", cfg.EncryptionKey)
	assert.Equal(t, "/default-tls-cert", cfg.TLSCertPath)
	assert.Equal(t, "/default-tls-key", cfg.TLSKeyPath)
	assert.Equal(t, 30, cfg.ShutdownTimeout)
	assert.Equal(t, "info", cfg.LoggerLvl)
}

func TestGetConfig_EnvironmentVariables(t *testing.T) {
	resetFlags()
	os.Args = []string{
		"cmd",
	}

	err := os.Setenv("ADDRESS", "localhost:8081")
	require.NoError(t, err, "unexpected error setting ADDRESS")
	err = os.Setenv("DATABASE_DSN", "env-dsn")
	require.NoError(t, err, "unexpected error setting DATABASE_DSN")
	err = os.Setenv("FILE_STORAGE_PATH", "/env-storage-path")
	require.NoError(t, err, "unexpected error setting FILE_STORAGE_PATH")
	err = os.Setenv("AUTH_KEY", "env-auth-key")
	require.NoError(t, err, "unexpected error setting AUTH_KEY")
	err = os.Setenv("ENCRYPTION_KEY", "env-encrypt-key")
	require.NoError(t, err, "unexpected error setting ENCRYPTION_KEY")
	err = os.Setenv("TLS_CERT", "/env-tls-cert")
	require.NoError(t, err, "unexpected error setting TLS_CERT")
	err = os.Setenv("TLS_KEY", "/env-tls-key")
	require.NoError(t, err, "unexpected error setting TLS_KEY")

	defer func() {
		os.Clearenv()
	}()

	defaultCfg := config.ServerConfig{}
	cfg, err := GetConfig(defaultCfg)
	require.NoError(t, err, "unexpected error getting config")

	assert.Equal(t, "localhost:8081", cfg.URL)
	assert.Equal(t, "env-dsn", cfg.DatabaseDSN)
	assert.Equal(t, "/env-storage-path", cfg.FileStoragePath)
	assert.Equal(t, "env-auth-key", cfg.AuthenticationKey)
	assert.Equal(t, "env-encrypt-key", cfg.EncryptionKey)
	assert.Equal(t, "/env-tls-cert", cfg.TLSCertPath)
	assert.Equal(t, "/env-tls-key", cfg.TLSKeyPath)
}

func TestGetConfig_CommandLineFlags(t *testing.T) {
	resetFlags()

	os.Args = []string{
		"cmd",
		"-h=localhost:8083",
		"-d=cmd-dsn",
		"-f=/cmd-storage-path",
		"-a=cmd-auth-key",
		"-e=cmd-encrypt-key",
		"-tls-cert=/cmd-tls-cert",
		"-tls-key=/cmd-tls-key",
	}

	defaultCfg := config.ServerConfig{}
	cfg, err := GetConfig(defaultCfg)
	require.NoError(t, err, "unexpected error getting config")

	assert.Equal(t, "localhost:8083", cfg.URL)
	assert.Equal(t, "cmd-dsn", cfg.DatabaseDSN)
	assert.Equal(t, "/cmd-storage-path", cfg.FileStoragePath)
	assert.Equal(t, "cmd-auth-key", cfg.AuthenticationKey)
	assert.Equal(t, "cmd-encrypt-key", cfg.EncryptionKey)
	assert.Equal(t, "/cmd-tls-cert", cfg.TLSCertPath)
	assert.Equal(t, "/cmd-tls-key", cfg.TLSKeyPath)
}

func TestGetConfig_JSONConfigFile(t *testing.T) {
	resetFlags()

	jsonCfg := `
	{
		"address": "localhost:8084",
		"database_dsn": "json-dsn",
		"file_storage_directory": "/json-storage-path",
		"authentication_key": "json-auth-key",
		"encryption_key": "json-encrypt-key",
		"tls_cert": "/json-tls-cert",
		"tls_key": "/json-tls-key",
		"shutdown_timeout": 90,
		"logger_level": "warn"
	}`
	jsonCfgFile, err := os.CreateTemp("", "config*.json")
	require.NoError(t, err, "unexpected error creating json config file")
	defer os.Remove(jsonCfgFile.Name())

	_, err = jsonCfgFile.Write([]byte(jsonCfg))
	require.NoError(t, err, "unexpected error writing json config file")
	err = jsonCfgFile.Close()
	require.NoError(t, err, "unexpected error closing json config file")

	os.Args = []string{"cmd", "-c", jsonCfgFile.Name()}

	defaultCfg := config.ServerConfig{}
	cfg, err := GetConfig(defaultCfg)
	require.NoError(t, err, "unexpected error getting config")

	assert.Equal(t, "localhost:8084", cfg.URL)
	assert.Equal(t, "json-dsn", cfg.DatabaseDSN)
	assert.Equal(t, "/json-storage-path", cfg.FileStoragePath)
	assert.Equal(t, "json-auth-key", cfg.AuthenticationKey)
	assert.Equal(t, "json-encrypt-key", cfg.EncryptionKey)
	assert.Equal(t, "/json-tls-cert", cfg.TLSCertPath)
	assert.Equal(t, "/json-tls-key", cfg.TLSKeyPath)
	assert.Equal(t, 90, cfg.ShutdownTimeout)
	assert.Equal(t, "warn", cfg.LoggerLvl)
}

func resetFlags() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
}
