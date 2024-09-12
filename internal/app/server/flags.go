package server

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/caarlos0/env"

	config "github.com/Stern-Ritter/gophkeeper/internal/config/server"
	"github.com/Stern-Ritter/gophkeeper/internal/utils"
)

type appConfig struct {
	URL               string `json:"address,omitempty"`
	DatabaseDSN       string `json:"database_dsn,omitempty"`
	FileStoragePath   string `json:"file_storage_directory,omitempty"`
	AuthenticationKey string `json:"authentication_key,omitempty"`
	EncryptionKey     string `json:"encryption_key,omitempty"`
	TLSCertPath       string `json:"tls_cert,omitempty"`
	TLSKeyPath        string `json:"tls_key,omitempty"`
	ShutdownTimeout   int    `json:"shutdown_timeout,omitempty"`
	LoggerLvl         string `json:"logger_level,omitempty"`
}

// GetConfig initializes the server config by parsing command-line flags, environment variables, and a JSON config file.
// It returns the initialized server config and any parsing error encountered.
//
// The priority of configuration values is as follows (from highest to lowest):
// 1. Environment variables
// 2. Command-line flags
// 3. JSON config file
// 4. Default config
func GetConfig(defaultCfg config.ServerConfig) (config.ServerConfig, error) {
	cfg := config.ServerConfig{}

	parseFlags(&cfg)

	err := env.Parse(&cfg)
	if err != nil {
		return cfg, err
	}

	cfgFile := strings.TrimSpace(cfg.ConfigFile)
	needParseJSONConfig := len(cfgFile) > 0
	if needParseJSONConfig {
		err = parseJSONConfig(&cfg, cfgFile)
		if err != nil {
			return cfg, err
		}
	}

	mergeDefaultConfig(&cfg, defaultCfg)

	trimStringVarsSpaces(&cfg)

	return cfg, nil
}

func parseFlags(cfg *config.ServerConfig) {
	flag.StringVar(&cfg.URL, "h", "", "address and port to run server in format <host>:<port>")
	flag.StringVar(&cfg.DatabaseDSN, "d", "", "database dsn")
	flag.StringVar(&cfg.FileStoragePath, "f", "", "metrics storage file path")
	flag.StringVar(&cfg.AuthenticationKey, "a", "", "secret authentication key")
	flag.StringVar(&cfg.EncryptionKey, "e", "", "encryption key")
	flag.StringVar(&cfg.TLSCertPath, "tls-cert", "", "path to tls certificate")
	flag.StringVar(&cfg.TLSKeyPath, "tls-key", "", "path to tls key")
	flag.StringVar(&cfg.ConfigFile, "c", "", "path to json config file")
	flag.Parse()
}

func parseJSONConfig(cfg *config.ServerConfig, fPath string) error {
	data, err := os.ReadFile(fPath)
	if err != nil {
		return fmt.Errorf("read config file %s: %w", fPath, err)
	}

	appCfg := appConfig{}
	err = json.Unmarshal(data, &appCfg)
	if err != nil {
		return fmt.Errorf("parse config file %s: %w", fPath, err)
	}

	mergeJSONConfig(cfg, appCfg)
	return nil
}

func mergeJSONConfig(cfg *config.ServerConfig, appCfg appConfig) {
	cfg.URL = utils.Coalesce(cfg.URL, appCfg.URL)
	cfg.DatabaseDSN = utils.Coalesce(cfg.DatabaseDSN, appCfg.DatabaseDSN)
	cfg.FileStoragePath = utils.Coalesce(cfg.FileStoragePath, appCfg.FileStoragePath)
	cfg.AuthenticationKey = utils.Coalesce(cfg.AuthenticationKey, appCfg.AuthenticationKey)
	cfg.EncryptionKey = utils.Coalesce(cfg.EncryptionKey, appCfg.EncryptionKey)
	cfg.TLSCertPath = utils.Coalesce(cfg.TLSCertPath, appCfg.TLSCertPath)
	cfg.TLSKeyPath = utils.Coalesce(cfg.TLSKeyPath, appCfg.TLSKeyPath)
	cfg.ShutdownTimeout = utils.Coalesce(cfg.ShutdownTimeout, appCfg.ShutdownTimeout)
	cfg.LoggerLvl = utils.Coalesce(cfg.LoggerLvl, appCfg.LoggerLvl)
}

func mergeDefaultConfig(cfg *config.ServerConfig, defaultCfg config.ServerConfig) {
	cfg.URL = utils.Coalesce(cfg.URL, defaultCfg.URL)
	cfg.DatabaseDSN = utils.Coalesce(cfg.DatabaseDSN, defaultCfg.DatabaseDSN)
	cfg.FileStoragePath = utils.Coalesce(cfg.FileStoragePath, defaultCfg.FileStoragePath)
	cfg.AuthenticationKey = utils.Coalesce(cfg.AuthenticationKey, defaultCfg.AuthenticationKey)
	cfg.EncryptionKey = utils.Coalesce(cfg.EncryptionKey, defaultCfg.EncryptionKey)
	cfg.TLSCertPath = utils.Coalesce(cfg.TLSCertPath, defaultCfg.TLSCertPath)
	cfg.TLSKeyPath = utils.Coalesce(cfg.TLSKeyPath, defaultCfg.TLSKeyPath)
	cfg.ConfigFile = utils.Coalesce(cfg.ConfigFile, defaultCfg.ConfigFile)
	cfg.ShutdownTimeout = utils.Coalesce(cfg.ShutdownTimeout, defaultCfg.ShutdownTimeout)
	cfg.LoggerLvl = utils.Coalesce(cfg.LoggerLvl, defaultCfg.LoggerLvl)
}

func trimStringVarsSpaces(cfg *config.ServerConfig) {
	cfg.URL = strings.TrimSpace(cfg.URL)
	cfg.DatabaseDSN = strings.TrimSpace(cfg.DatabaseDSN)
	cfg.FileStoragePath = strings.TrimSpace(cfg.FileStoragePath)
	cfg.AuthenticationKey = strings.TrimSpace(cfg.AuthenticationKey)
	cfg.EncryptionKey = strings.TrimSpace(cfg.EncryptionKey)
	cfg.TLSCertPath = strings.TrimSpace(cfg.TLSCertPath)
	cfg.TLSKeyPath = strings.TrimSpace(cfg.TLSKeyPath)
	cfg.ConfigFile = strings.TrimSpace(cfg.ConfigFile)
	cfg.LoggerLvl = strings.TrimSpace(cfg.LoggerLvl)
}
