package rabbitmq

import (
	"github.com/streadway/amqp"
)

type Config struct {
	Host        string
	Port        string
	UserName    string
	Password    string
	RetryTTL    int
	Concurrency int
}

type Connection struct {
	connection *amqp.Connection
	config     *Config
	service    string
}
