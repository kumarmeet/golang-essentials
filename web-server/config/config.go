package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string
	AppJWT  string
	AppEnv  string
	AppDB   string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load(fmt.Sprintf(".env.%s", os.Getenv("APP_ENV")))

	fmt.Println(fmt.Sprintf(".env.%s", os.Getenv("APP_ENV")))

	if err != nil {
		return nil, fmt.Errorf("Error loading .%s.env file", os.Getenv("APP_ENV"))
	}

	return &Config{
		AppPort: os.Getenv("APP_PORT"),
		AppJWT:  os.Getenv("JWT_SECRETKEY"),
		AppEnv:  os.Getenv("APP_ENV"),
		AppDB:   os.Getenv("DB_NAME"),
	}, nil
}
