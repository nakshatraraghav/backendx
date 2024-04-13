package lib

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvSchema struct {
	Addr       string `validate:"required"`
	Envionment string `validate:"required"`
}

var ENV EnvSchema

var logger = GetLogger()

func LoadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		logger.Fatal().Str("message", "failed to find the .env.local file please check for it again").Err(err).Send()
	}

	e := EnvSchema{}

	e.Addr = os.Getenv("ADDR")
	e.Envionment = os.Getenv("ENVIRONMENT")

	err = Validate.Struct(e)
	if err != nil {
		logger.Fatal().Str("message", "failed to validate the environment variables file please check it again").Err(err).Send()
	}

	ENV = e
}
