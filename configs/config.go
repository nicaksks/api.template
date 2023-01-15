package configs

import (
	"github.com/joho/godotenv"
)

func StartConfig() error {

	if err := godotenv.Load(); err != nil {
		return err
	}

	return nil
}
