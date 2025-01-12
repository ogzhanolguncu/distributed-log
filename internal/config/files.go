package config

import (
	"os"
	"path/filepath"
)

var (
	CAFile         string
	ServerCertFile string
	ServerKeyFile  string
)

func init() {
	CAFile = configFile("ca.pem")
	ServerCertFile = configFile("server.pem")
	ServerKeyFile = configFile("server-key.pem")

	// Ensure the config directory exists
	dir := filepath.Dir(CAFile)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		panic(err)
	}
}

func configFile(filename string) string {
	if dir := os.Getenv("CONFIG_DIR"); dir != "" {
		return filepath.Join(dir, filename)
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return filepath.Join(homeDir, ".proglog", filename)
}
