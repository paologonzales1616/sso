package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type envConfig struct {
	Port string
}

var Env *envConfig

func ReadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	Env = &envConfig{
		Port: getEnv("PORT", "3000")}
}

func getEnv(key string, defaultVal string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultVal
}
