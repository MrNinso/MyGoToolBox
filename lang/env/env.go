package env

import (
	"os"
	"strconv"
)

func GetEnv(key, defaultValue string) string {
	s := os.Getenv(key)
	if s == "" {
		return defaultValue
	} else {
		return s
	}
}

func SetEnv(key, value string) (hasUpdated bool, err error) {
	_, exists := os.LookupEnv(key)
	return exists, os.Setenv(key, value)
}

func GetIntFromEnv(key string, defaulValue int) int {
	v := GetEnv(key, "")

	i, err := strconv.Atoi(v)

	if err == nil {
		return i
	} else {
		return defaulValue
	}
}
