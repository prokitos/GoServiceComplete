package config

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

const projectDirName = "GoServiceComplete"

func LoadEnv(envName string) {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `\\internal\\config\\` + envName + `.env`)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
