package rabbit_utils

import (
	"os"
)

type rabbitmq struct {
	username string
	password string
	host     string
	port     string
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

func GetConnectionString() string {
	var rabbit_connection string

	rabbitConfig := rabbitmq{
		username: getEnv("RABBITMQ_USERNAME", "guest"),
		password: getEnv("RABBITMQ_PASSWORD", "guest"),
		host:     getEnv("RABBITMQ_HOST", "localhost"),
		port:     getEnv("RABBITMQ_PORT", "5672"),
	}
	rabbit_connection = "amqp://" + rabbitConfig.username + ":" + rabbitConfig.password + "@" + rabbitConfig.host + ":" + rabbitConfig.port + "/"
	return rabbit_connection
}
