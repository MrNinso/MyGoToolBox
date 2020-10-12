package env

import (
	. "github.com/MrNinso/MyGoToolBox/lang/ifs"
	"os"
	"strconv"
)

func GetEnv(key, defaultValue string) string {
	return IfStringEmpty(os.Getenv(key), defaultValue)
}

func SetEnv(key, value string) (hasUpdated bool, err error) {
	_, exists := os.LookupEnv(key)
	return exists, os.Setenv(key, value)
}

func GetIntFromEnv(key string, defaulValue int) int {
	v := GetEnv(key, "")

	i, err := strconv.Atoi(v)

	return IfReturn(err == nil, i, defaulValue).(int)
}
