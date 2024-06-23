package env

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvApp struct {
	API_PORT string

	MONGO_URL      string
	MONGO_DB_NAME  string
	MONGO_USERNAME string
	MONGO_PASSWORD string
}

func GetEnvApp(envFile string) EnvApp {
	if os.Getenv("ENV") != "PRO" {
		err := godotenv.Load(envFile)

		if err != nil {
			panic(err)
		}
	}

	return EnvApp{
		API_PORT:       os.Getenv("API_PORT"),
		MONGO_URL:      os.Getenv("MONGO_URL"),
		MONGO_DB_NAME:  os.Getenv("MONGO_DB_NAME"),
		MONGO_USERNAME: os.Getenv("MONGO_USERNAME"),
		MONGO_PASSWORD: os.Getenv("MONGO_PASSWORD"),
	}

}
