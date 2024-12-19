package utils

import (
	"fmt"
	"os"
)

func LoadEnvVal(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("%s not set in environment variables", key)
	}
	return value, nil
}
