package env

import (
	"fmt"
	"os"
	"strconv"
)

func GetDBURI() (string, error) {
	const envName = "POSTGRES_URI"

	uri := os.Getenv(envName)
	if uri == "" {
		return "", fmt.Errorf("%s not set in environment variables", envName)
	}

	return uri, nil
}

func GetAPIPort() (int, error) {
	const envName = "BACKEND_PORT"

	port, err := strconv.Atoi(os.Getenv(envName))
	if err != nil {
		return 0, fmt.Errorf("%s not set in environment variables", envName)
	}

	return port, nil
}
