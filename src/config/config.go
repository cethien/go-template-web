package config

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
)

var (
	AppName         string = "Go Web Template"
	DbUrl           string
	DbMigrationsUrl string
)

func IsDevelopment() bool {
	return os.Getenv("APP_ENV") == "development"
}

func GetHttpPort() int {
	portEnv, ok := os.LookupEnv("PORT")
	if ok {
		i, err := strconv.Atoi(portEnv)
		if err != nil {
			slog.Error(fmt.Sprintf("value '%v' of env 'PORT' is not valid", portEnv))
		} else {
			return i
		}
	}

	if IsDevelopment() {
		return 8080
	}

	return 80
}

func Init() error {
	var err error

	DbUrl, err = getEnv("DB_URL")
	if err != nil {
		return err
	}

	DbMigrationsUrl, err = getEnv("DB_MIGRATIONS_URL")
	if err != nil {
		return err
	}

	return nil
}

func getEnv(name string) (string, error) {
	env, ok := os.LookupEnv(name)
	if !ok {
		return "", fmt.Errorf("env '%v' not found", name)
	}
	return env, nil
}
