package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/gommon/log"
)

var (
	Host                 = loadEnvOrDefault("HOST", "0.0.0.0")
	Port                 = loadEnvOrDefault("PORT", "8080")
	BaseURL              = loadEnvOrDefault("BASE_URL", fmt.Sprintf("http://%s:%s", Host, Port))
	GoEnv                = env(loadEnvOrDefault("GO_ENV", "development"))
	LogLevel             = logLevel(loadEnvOrDefault("LOG_LEVEL", "INFO"))
	SMTPHost             = loadEnvOrDefault("SMTP_HOST", "localhost")
	SMTPPort             = loadIntEnvOrDefault("SMTP_PORT", 1025)
	SMTPFrom             = loadEnvOrDefault("SMTP_FROM", "test@localhost")
	SMTPPassword         = loadEnvOrDefault("SMTP_PASSWORD", "")
	SMTPSSL              = loadBoolOrDefault("SMTP_SSL", false)
	MongoHost            = loadEnvOrDefault("MONGO_HOST", "localhost")
	MongoPort            = loadEnvOrDefault("MONGO_PORT", "27017")
	MongoUser            = loadEnvOrDefault("MONGO_USER", "root")
	MongoPassword        = loadEnvOrDefault("MONGO_PASSWORD", "root")
	MongoDBName          = loadEnvOrDefault("MONGO_DB_NAME", "factsfood")
	MongoCtxTimeout      = time.Duration(loadIntEnvOrDefault("MONGO_CONTEXT_TIMEOUT_SECONDS", 10)) * time.Second
	ActivationExpiration = time.Duration(loadIntEnvOrDefault("ACTIVATION_EXPIRATION", 20)) * time.Minute
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
		fmt.Println("warning! invalid log level:", level)
		lvl = log.INFO
	}
	return lvl
}

func loadBoolOrDefault(key string, defaultValue bool) bool {
	value := loadEnvOrDefault(key, fmt.Sprint(defaultValue))
	b, err := strconv.ParseBool(value)
	if err != nil {
		fmt.Println("warning! invalid key value (bool conversion):", key)
		return defaultValue
	}
	return b
}

func loadIntEnvOrDefault(key string, defaultValue int) int {
	value := loadEnvOrDefault(key, fmt.Sprint(defaultValue))
	num, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println("warning! invalid key value (int conversion):", key)
		return defaultValue
	}
	return num
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
		fmt.Println("warning! invalid env type:", envType)
		return DEVELOPMENT

	}
}
