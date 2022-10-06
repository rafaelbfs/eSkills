package security

import (
	"github.com/joho/godotenv"
	"log"
)

const (
	GOOGLE_CLIENT_SECRET = "GOOGLE_OAUTH_CLIENT_SECRET"
	GOOGLE_CLIENT_ID     = "GOOGLE_OAUTH_CLIENT_ID"
	SESSION_SECRET       = "SESSION_SECRET"
)

var ConfigMap map[string]string

func CfgGet(key string) string {
	if len(ConfigMap) == 0 {
		panic("No configuration is loaded, call 'Initialize' function with the path to your .env file")
	}
	return ConfigMap[key]
}

func Initialize(envFilePath string) {
	var err error
	ConfigMap, err = godotenv.Read(envFilePath)

	if err != nil {
		log.Panicf("Unable to provision configuration due to %v", err)
	}
}
