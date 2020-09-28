package main

import (
	"os"
)

type configuration struct {
	apiPort    string
	dbEndpoint string
	dbUsername string
	dbPassword string
}

func getConfig() *configuration {
	config := &configuration{}
	config.apiPort = getEnv("API_PORT", ":8080")
	config.dbEndpoint = getEnv("DB_ENDPOINT", "127.0.0.1:3306")
	config.dbUsername = getEnv("DB_USERNAME", "root")
	config.dbPassword = getEnv("DB_PASSWORD", "password")
	return config
}

func getEnv(envVar, fallback string) string {
	if value, ok := os.LookupEnv(envVar); ok {
		return value
	}
	return fallback
}
