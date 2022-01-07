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

func getConnectionString() string {
	var rabbit_connection string

	rabbitConfig := rabbitmq{
		username: os.Getenv("RABBITMQ_USERNAME"),
		password: os.Getenv("RABBITMQ_PASSWORD"),
		host:     os.Getenv("RABBITMQ_HOST"),
		port:     os.Getenv("RABBITMQ_PORT"),
	}
	rabbit_connection = "amqp://" + rabbitConfig.username + ":" + rabbitConfig.password + "@" + rabbitConfig.host + ":" + rabbitConfig.port + "/"
	return rabbit_connection
}
