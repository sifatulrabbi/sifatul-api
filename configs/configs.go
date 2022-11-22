package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Configs struct {
	SMTP_USER     string
	SMTP_PASSWORD string
	PORT          string
}

func LoadConfigs() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	if v := os.Getenv("SMTP_USER"); v == "" {
		log.Fatal("SMTP_USER not found")
	}
	if v := os.Getenv("SMTP_PASSWORD"); v == "" {
		log.Fatal("SMTP_PASSWORD not found")
	}
	if v := os.Getenv("PORT"); v == "" {
		log.Fatal("PORT not found")
	}
}

func GetConfigs() *Configs {
	c := &Configs{}
	c.PORT = os.Getenv("PORT")
	c.SMTP_PASSWORD = os.Getenv("SMTP_PASSWORD")
	c.SMTP_USER = os.Getenv("SMTP_USER")
	return c
}
