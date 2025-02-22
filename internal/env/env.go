package env

import (
	"os"
	"strconv"
)

func GetString(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}

func GetInt(key string, fallback int) int {
	valueAsInt, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}

	value, err := strconv.Atoi(valueAsInt)
	if err != nil {
		return fallback
	}

	return value
}
