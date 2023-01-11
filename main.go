package rabbitmq

import (
	"fmt"
	"github.com/ihatiko/log"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"time"
)

func (connection *Connection) GetExchangeName(publisherKey, entityKey string) string {
	return fmt.Sprintf(`%s_%s`, publisherKey, entityKey)
}

func (connection *Connection) GetQueueName(consumerKey, routingKey string) string {
	return fmt.Sprintf(`%s_%s_Queue`, consumerKey, routingKey)
}

func NewConnection(serviceName string, config *Config) (*Connection, error) {
	connection := getAmqpConnection(config)
	connectRabbitMQ, err := amqp.Dial(connection)
	if err != nil {
		return &Connection{}, errors.Wrap(err, "Amqp.Dial")
	}

	dialedConnection := Connection{
		connection: connectRabbitMQ,
		config:     config,
		service:    serviceName,
	}
	go func() {
		for {
			<-connectRabbitMQ.NotifyClose(make(chan *amqp.Error))
			for i := 0; i < 5; i++ {
				if err := dialedConnection.connect(); err != nil {
					log.Error(err.Error())
				}
				time.Sleep(1000)
			}
			log.Fatal("TimeOutError rabbitmq dial")
		}
	}()
	return &dialedConnection, nil
}

var buffer = map[string]*Consumer{}

func (connection *Connection) connect() error {
	connection, err := NewConnection(connection.service, connection.config)
	if err != nil {
		return err
	}
	for _, value := range buffer {
		consumer, err := connection.GetConsumer(value.Exchange, value.queueName, value.routerKey)
		if err != nil {
			return err
		}
		go func() {
			_, err := consumer.Consume(consumer.handler)
			if err != nil {
				log.Error(err.Error())
			}
		}()
	}

	return nil
}

func getAmqpConnection(config *Config) string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/", config.UserName, config.Password, config.Host, config.Port)
}

func (connection *Connection) Close() {
	err := connection.connection.Close()
	if err != nil {
		log.Error()
	}
}
