package config

import (
	"github.com/labstack/gommon/log"
	"os"
	"strings"
)

var (
	Host     = loadEnvOrDefault("HOST", "0.0.0.0")
	Port     = loadEnvOrDefault("PORT", "8080")
	GoEnv    = env(loadEnvOrDefault("GO_ENV", "development"))
	LogLevel = logLevel(loadEnvOrDefault("LOG_LEVEL", "INFO"))
)

type EnvType uint8

const (
	DEVELOPMENT EnvType = iota
	TEST
	PRODUCTION
)

func loadEnvOrDefault(key string, defaultValue string) string {
	value, exist := os.LookupEnv(key)
	if !exist {
		return defaultValue
	} else {
		return value
	}
}

func logLevel(level string) log.Lvl {
	var lvl log.Lvl
	switch strings.ToUpper(level) {
	case "DEBUG":
		lvl = log.DEBUG
	case "INFO":
		lvl = log.INFO
	case "WARN":
		lvl = log.WARN
	case "ERROR":
		lvl = log.ERROR
	case "OFF":
		lvl = log.OFF
	default:
		println("warning! invalid log level:", level)
		lvl = log.INFO
	}
	return lvl

}

func env(envType string) EnvType {
	switch strings.ToUpper(envType) {
	case "DEVELOPMENT":
		return DEVELOPMENT
	case "TEST":
		return TEST
	case "PRODUCTION":
		return PRODUCTION
	default:
		println("warning! invalid env type:", envType)
		return DEVELOPMENT

	}

}
