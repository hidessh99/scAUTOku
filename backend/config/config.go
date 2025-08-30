package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds the application configuration
type Config struct {
	APIKey     string
	JWTSecret  string
	AdminUser  string
	AdminPass  string
	ServerPort string
	ServerHost string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

// AppConfig holds the global application configuration
var AppConfig *Config

// LoadConfig loads configuration from environment variables
func LoadConfig() {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	AppConfig = &Config{
		APIKey:     getEnv("API_KEY", "default_api_key"),
		JWTSecret:  getEnv("JWT_SECRET", "default_jwt_secret"),
		AdminUser:  getEnv("ADMIN_USERNAME", "admin"),
		AdminPass:  getEnv("ADMIN_PASSWORD", "admin"),
		ServerPort: getEnv("PORT", "3000"),
		ServerHost: getEnv("HOST", "localhost"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "vpnuser"),
		DBPassword: getEnv("DB_PASSWORD", "vpnpassword"),
		DBName:     getEnv("DB_NAME", "vpnaccounts"),
	}
}

// getEnv retrieves environment variable or returns default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// getEnvAsInt retrieves environment variable as integer or returns default value
func getEnvAsInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
