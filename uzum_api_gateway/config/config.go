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
	Environment string
	Version     string

	// RedisHost     string
	// RedisPort     int
	// RedisPassword string

	CatalogServiceHost string
	CatalogServicePort string

	UserServiceHost string
	UserServicePort string
	
	OrderServiceHost string
	OrderServicePort string

	LogLevel string
	HTTPPort string
}

// Load loads environment vars and inflates Config
func Load() Config {

	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("No .env file found")
	}
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", DebugMode))
	c.Version = cast.ToString(getOrReturnDefault("VERSION", "1.0"))

	// c.RedisHost = cast.ToString(getOrReturnDefault("REDIS_HOST", "127.0.0.1"))
	// c.RedisPort = cast.ToInt(getOrReturnDefault("REDIS_PORT", 6379))
	// c.RedisPassword = cast.ToString(getOrReturnDefault("REDIS_PASSWORD", "3EEdwhDOfx"))

	c.CatalogServiceHost = cast.ToString(getOrReturnDefault("CATALOG_SERVICE_HOST", "localhost"))
	c.CatalogServicePort = cast.ToString(getOrReturnDefault("CATALOG_GRPC_PORT", "8080"))

	c.UserServiceHost = cast.ToString(getOrReturnDefault("USER_SERVICE_HOST", "localhost"))
	c.UserServicePort = cast.ToString(getOrReturnDefault("USER_SEVICE_PORT", "8081"))
	
	c.OrderServiceHost = cast.ToString(getOrReturnDefault("ORDER_SERVICE_HOST", "localhost"))
	c.OrderServicePort = cast.ToString(getOrReturnDefault("ORDER_SERVICE_PORT", "8082"))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":1234"))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	if os.Getenv(key) == "" {
		return defaultValue
	}

	return os.Getenv(key)
}
