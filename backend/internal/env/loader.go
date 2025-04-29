package env

import (
	"fmt"
	"os"
	"strconv"
)

func getEnvAsString(envName string) (string, error) {
	value := os.Getenv(envName)
	if value == "" {
		return "", fmt.Errorf("%s not set in environment variables", envName)
	}
	return value, nil
}

func GetDBURI() (string, error) {
	return getEnvAsString("POSTGRES_URI")
}

func GetAPIPort() (int, error) {
	const envName = "BACKEND_PORT"

	port, err := strconv.Atoi(os.Getenv(envName))
	if err != nil {
		return 0, fmt.Errorf("%s not set in environment variables", envName)
	}

	return port, nil
}

func GetFrontendURI() (string, error) {
	return getEnvAsString("FRONTEND_URI")
}
