package config

import "os"

var (
	Host = loadEnvOrDefault("HOST", "0.0.0.0")
	Port = loadEnvOrDefault("PORT", "8080")
)

func loadEnvOrDefault(key string, defaultValue string) string {
	value, exist := os.LookupEnv(key)
	if !exist {
		return defaultValue
	} else {
		return value
	}
}
