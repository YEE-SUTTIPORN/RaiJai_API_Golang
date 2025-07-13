package config

import (
	"os"
)

func GetDSN() string {
    return os.Getenv("DATABASE_RAIJAI")
}