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
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln(".env file not found")
	}

	switch {
	case os.Getenv("DATABASE_URL") == "":
		log.Fatalln("No DATABASE_URL found")
	case os.Getenv("PORT") == "":
		log.Fatalln("No PORT found")
	case os.Getenv("JWT_SECRET") == "":
		log.Fatalln("No JWT_SECRET found")
	default:
		log.Println("Complete env entries")
	}

	return &Config{
		DatabaseUrl: os.Getenv("DATABASE_URL"),
		JwtSecret: os.Getenv("JWT_SECRET"),
		Port: os.Getenv("PORT"),
	}
}
