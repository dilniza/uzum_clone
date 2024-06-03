package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const (
	// DebugMode indicates service mode is debug.
	DebugMode = "debug"
	// TestMode indicates service mode is test.
	TestMode = "test"
	// ReleaseMode indicates service mode is release.
	ReleaseMode = "release"
)

// Config ...
type Config struct {
	ServiceName string
	Environment string // develop, staging, production
	Version     string

	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string
	PostgresMaxConnections int32

	// RedisHost     string
	// RedisPort     int
	// RedisPassword string

	CatalogServiceHost string
	CatalogServicePort string

	LogLevel string
	HTTPPort string
}

// Load loads environment vars and inflates Config
func Load() Config {

	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("No .env file found")
	}
	c := Config{}

	c.ServiceName = cast.ToString(getOrReturnDefaultValue("SERVICE_NAME", "catalog_service"))
	c.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", DebugMode))
	c.Version = cast.ToString(getOrReturnDefaultValue("VERSION", "1.0"))

	c.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 8082))
	c.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "admin"))
	c.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "admin"))
	c.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "project"))
	c.PostgresMaxConnections = cast.ToInt32(getOrReturnDefaultValue("POSTGRES_MAX_CONNECTIONS", 30))

	c.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":8080"))
	//c.RedisHost = cast.ToString(getOrReturnDefault("REDIS_HOST", "127.0.0.1"))
	//c.RedisPort = cast.ToInt(getOrReturnDefault("REDIS_PORT", 6379))
	//c.RedisPassword = cast.ToString(getOrReturnDefault("REDIS_PASSWORD", ""))

	c.CatalogServiceHost = cast.ToString(getOrReturnDefaultValue("CATALOG_SERVICE_HOST", "localhost"))
	c.CatalogServicePort = cast.ToString(getOrReturnDefaultValue("CATALOG_GRPC_PORT", "8082"))

	return c
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	if os.Getenv(key) == "" {
		return defaultValue
	}

	return os.Getenv(key)
}
