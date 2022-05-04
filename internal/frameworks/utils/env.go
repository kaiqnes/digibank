package utils

import "os"

func GetEnvOrDefault(key, defaultValue string) string {
	envVar := os.Getenv(key)
	if envVar == "" {
		return defaultValue
	}
	return envVar
}
