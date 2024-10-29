package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	PrivateAddr    string
	PublicAddr     string
	AppVersion     string
	AppId          string
	AppName        string
	DebugMode      bool
}

func NewAppConfig(version string) *AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Can't load .env file : %s\n", err.Error())
	}

	appId, err := os.Hostname()
	if err != nil {
		log.Printf("Can't get os hostname : %s\n", err.Error())
	}

	return &AppConfig{
		PrivateAddr:    os.Getenv("PRIVATE_ADDR"),
		PublicAddr:     os.Getenv("PUBLIC_ADDR"),
		AppVersion:     version,
		AppId:          appId,
		AppName:        "go-microservice",
		DebugMode:      os.Getenv("MODE") == "DEBUG",
	}
}