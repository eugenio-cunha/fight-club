package env

import (
	"fmt"
	"os"
)

func GetEnv(key string) string {
	value := os.Getenv(key)

	if value == "" {
		panic(fmt.Errorf("Invalid/Missing environment variable %s", key))
	}

	return value
}