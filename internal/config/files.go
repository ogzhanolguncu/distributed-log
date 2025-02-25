package config

import (
	"os"
	"path/filepath"
)

var (
	CAFile               = configFile("ca.pem")
	ServerCertFile       = configFile("server.pem")
	ServerKeyFile        = configFile("server-key.pem")
	ClientCertFile       = configFile("client.pem")
	ClientKeyFile        = configFile("client-key.pem")
	RootClientCertFile   = configFile("root-client.pem")
	RootClientKeyFile    = configFile("root-client-key.pem")
	NobodyClientCertFile = configFile("nobody-client.pem")
	NobodyClientKeyFile  = configFile("nobody-client-key.pem")
	ACLModelFile         = configFile("model.conf")
	ACLPolicyFile        = configFile("policy.csv")
)

// func init() {
// 	CAFile = configFile("ca.pem")
// 	ServerCertFile = configFile("server.pem")
// 	ServerKeyFile = configFile("server-key.pem")
// 	ClientCertFile = configFile("client.pem")
// 	ClientKeyFile = configFile("client-key.pem")
//
// 	// Ensure the config directory exists
// 	dir := filepath.Dir(CAFile)
// 	err := os.MkdirAll(dir, 0755)
// 	if err != nil {
// 		panic(err)
// 	}
// }

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
