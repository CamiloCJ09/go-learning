package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Get(key string) string {
	return os.Getenv(key)
}

func (c *Config) Set(key, value string) {
	os.Setenv(key, value)
}

func LoadConfig() {

	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

}
