package constant

import (
	"os"
	"sync"
)

// Config holds the application configuration values
type Config struct {
	DatabaseURL   string
	Port          string
	AdminUsername string
	AdminPassword string
}

// cfg holds the instance of Config to ensure it is initialized only once.
var cfg *Config
var once sync.Once

// InitConfig initializes the configuration object if it has not been initialized already.
func InitConfig() {
	once.Do(func() {
		cfg = &Config{
			DatabaseURL:   getEnv("DATABASE_URL", "host=localhost port=5432 user=postgres password=postgres dbname=ktaxes sslmode=disable"),
			Port:          getEnv("PORT", "8080"),
			AdminUsername: getEnv("ADMIN_USERNAME", "adminTax"),
			AdminPassword: getEnv("ADMIN_PASSWORD", "admin!"),
		}
	})
}

// Get returns the global configuration instance.
func Get() *Config {
	if cfg == nil {
		InitConfig()
	}
	return cfg
}

// getEnv tries to get a value from the environment and falls back to a default value if not found.
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
