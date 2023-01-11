package rabbitmq

import "github.com/streadway/amqp"

type Channel struct {
	ch            *amqp.Channel
	deadlyChannel *amqp.Channel
}
