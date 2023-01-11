package rabbitmq

import (
	"context"
	"fmt"
	"github.com/ihatiko/log"
	"github.com/streadway/amqp"
)

type Consumer struct {
	config     *Config
	connection *Connection
	Exchange   string
	queueName  string
	routerKey  string
	handler    func(context context.Context, d amqp.Delivery) bool
}

func (connection *Connection) GetConsumer(exchange, queueName, routerKey string) (*Consumer, error) {
	return &Consumer{
		connection: connection,
		config:     connection.config,
		Exchange:   exchange,
		queueName:  queueName,
		routerKey:  routerKey,
	}, nil
}

func (consumer *Consumer) Consume(handler func(context context.Context, d amqp.Delivery) bool) (*Channel, error) {
	deadlyExchange := fmt.Sprintf(`%s_%s_%s`, deadlyExchangePrefix, consumer.Exchange, consumer.connection.service)
	ch, err := consumer.connection.connection.Channel()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	buffer[fmt.Sprintf(`%s_%s_%s`, consumer.Exchange, consumer.routerKey, consumer.queueName)] = consumer
	_, err = ch.QueueDeclare(consumer.queueName, true, false, false, false, amqp.Table{
		xDeadLetterExchange: deadlyExchange,
	})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	err = ch.QueueBind(consumer.queueName, consumer.routerKey, consumer.Exchange, false, nil)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	err = consumer.initDeadlyExchange(ch, deadlyExchange, consumer.queueName, consumer.routerKey)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	prefetchCount := consumer.config.Concurrency * 4
	err = ch.Qos(prefetchCount, 0, false)
	if err != nil {
		return nil, err
	}

	messages, err := ch.Consume(
		consumer.queueName, // queue
		consumer.connection.service,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	go consumer.processMessage(messages, handler)
	return &Channel{ch: ch}, nil
}

func (consumer *Consumer) processMessage(messages <-chan amqp.Delivery,
	handler func(ctx context.Context, d amqp.Delivery) bool) {
	for i := 0; i < consumer.config.Concurrency; i++ {
		for message := range messages {
			ctx := context.Background()
			spCtx, _ := amqptracer.Extract(message.Headers)
			sp := opentracing.StartSpan(
				"ConsumeMessage",
				opentracing.FollowsFrom(spCtx),
			)
			ctx = opentracing.ContextWithSpan(ctx, sp)
			if handler(ctx, message) {
				err := message.Ack(false)
				if err != nil {
					log.FatalF("error Ack rabbitMq %v", err)
				}
			} else {
				err := message.Reject(false)
				if err != nil {
					log.FatalF("error Nack rabbitMq %v", err)
				}
			}
			sp.Finish()
		}
		log.Warn("Rabbit consumer closed")
	}
}

func (consumer *Consumer) initDeadlyExchange(ch *amqp.Channel, deadlyExchange, queueName, routerKey string) error {
	err := ch.ExchangeDeclare(
		deadlyExchange,
		defaultExchangeType,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Error("Rabbit consumer closed - critical Error")
		return err
	}
	innerEx := fmt.Sprintf(`%s_%s_%s`, innerExchange, deadlyExchangePrefix, queueName)
	err = ch.ExchangeDeclare(
		innerEx,
		defaultExchangeType,
		false,
		false,
		false,
		false,
		nil,
	)
	deadlyQueue := fmt.Sprintf(`%s_%s`, deadlyExchangePrefix, queueName)
	_, err = ch.QueueDeclare(
		deadlyQueue,
		true,
		false,
		false,
		false,
		amqp.Table{
			xDeadLetterExchange: innerEx,
			xDeadMessageTTL:     consumer.config.RetryTTL * 1000,
		},
	)

	if err != nil {
		return err
	}
	err = ch.QueueBind(consumer.queueName, routerKey, innerEx, false, nil)
	if err != nil {
		return err
	}
	return ch.QueueBind(deadlyQueue, routerKey, deadlyExchange, false, nil)
}
