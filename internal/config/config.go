package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseUrl string
	Port        string
	JwtSecret   string
	GoogleClientId string
	GoogleClientSecret string
	GoogleSecurityKey string
	GoogleRedirectUrl string
	AdminKey string
	AdminEmail string
}

func requireEnv(variable string) string {
	if os.Getenv(variable) == "" {
		log.Fatalf("%s is required in .env", variable)
		return ""
	}else{
		return os.Getenv(variable)
	}
}

func getEnv(variable string, defaultEnv string) string {
	if os.Getenv(variable) == "" {
		return defaultEnv
	}else{
		return os.Getenv(variable)
	}
}

func LoadConfig() *Config {
	_ = godotenv.Load(".env")

	ConfigData := map[string]string{
		"DATABASE_URL" : requireEnv("DATABASE_URL"),
		"PORT": getEnv("PORT", ":3333"),
		"JWT_SECRET": requireEnv("JWT_SECRET"),
		"GOOGLE_CLIENT_ID": requireEnv("GOOGLE_CLIENT_ID"),
		"GOOGLE_CLIENT_SECRET": requireEnv("GOOGLE_CLIENT_SECRET"),
		"GOOGLE_SECURITY_KEY": requireEnv("GOOGLE_SECURITY_KEY"),
		"GOOGLE_REDIRECT_URL": getEnv("GOOGLE_REDIRECT_URL", "http://localhost:3333/api/v1/auth/callback/google"),
		"ADMIN_KEY": requireEnv("ADMIN_KEY"),
		"ADMIN_EMAIL": requireEnv("ADMIN_EMAIL"),
	}

	return &Config{
		DatabaseUrl: ConfigData["DATABASE_URL"],
		JwtSecret: ConfigData["JWT_SECRET"],
		Port: ConfigData["PORT"],
		GoogleClientId: ConfigData["GOOGLE_CLIENT_ID"],
		GoogleClientSecret: ConfigData["GOOGLE_CLIENT_SECRET"],
		GoogleSecurityKey: ConfigData["GOOGLE_SECURITY_KEY"],
		GoogleRedirectUrl: ConfigData["GOOGLE_REDIRECT_URL"],
		AdminKey: ConfigData["ADMIN_KEY"],
		AdminEmail: ConfigData["ADMIN_EMAIL"],
	}
}
