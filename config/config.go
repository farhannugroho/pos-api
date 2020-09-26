package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Configuration struct {
	PageSize  int
	JwtSecret string

	// Image
	ImageSavePath string
	ImageMaxSize  int
	ImageAllowExt []string

	// Database
	DbHost     string
	DbUser     string
	DbPassword string
	DbName     string
	DbPort     string
	DbSSLMode  string
	DbTimeZone string
}

var Config = &Configuration{}

func Setup() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("config.Setup, Error loading .env file: %v", err)
	}

	Config.DbHost = os.Getenv("DB_HOST")
	Config.DbUser = os.Getenv("DB_USER")
	Config.DbPassword = os.Getenv("DB_PASS")
	Config.DbName = os.Getenv("DB_NAME")
	Config.DbPort = os.Getenv("DB_PORT")
	Config.DbSSLMode = os.Getenv("DB_SSLMODE")
	Config.DbTimeZone = os.Getenv("DB_TIMEZONE")
}
