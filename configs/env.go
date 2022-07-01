package configs

import (
	"log"
	"os"

	. "flaq.club/backend/configs/vars"
	"github.com/joho/godotenv"
)

func LoadEnv(env string) string {
	// Declare all the expected envs here

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading env")
	}

	envValues := map[string]string{}

	for _, env := range Envs {
		envValues[env] = os.Getenv(env)
	}

	return envValues[env]
}
