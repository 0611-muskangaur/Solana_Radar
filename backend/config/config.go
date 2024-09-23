// config/config.go
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() { //This function loads environment variables from a .env file into the application's environment using the godotenv.Load() method.
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file") //If the .env file cannot be loaded (e.g., if it doesn't exist or there is an error), it logs a fatal error and stops the application using log.Fatalf().
	}
}

func GetEnv(key string) string { //This function retrieves the value of an environment variable by its key using os.Getenv(). It returns the value as a string.
	return os.Getenv(key)
}
