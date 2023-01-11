package rabbitmq

import (
	"context"
	"encoding/json"
	"github.com/ihatiko/log"
	"github.com/streadway/amqp"
	"google.golang.org/protobuf/proto"
)

type Producer struct {
	channel  *amqp.Channel
	config   *Config
	exchange string
}

func (connection *Connection) GetCustomProducer(exchange, kind string) (*Producer, error) {
	ch, err := connection.connection.Channel()
	if err != nil {
		return &Producer{}, err
	}

	err = ch.ExchangeDeclare(
		exchange,
		kind,
		false,
		false,
		false,
		false,
		nil,
	)

	return &Producer{
		channel:  ch,
		config:   connection.config,
		exchange: exchange,
	}, err
}

func (connection *Connection) GetProducer(exchangeName string) (*Producer, error) {
	return connection.GetCustomProducer(exchangeName, defaultExchangeType)
}

func (producer *Producer) ProduceJson(ctx context.Context, routerKey string, data interface{}) error {
	sp := opentracing.SpanFromContext(ctx)
	defer func() {
		if sp != nil {
			sp.Finish()
		}
	}()
	body, err := json.Marshal(&data)
	if err != nil {
		return err
	}
	log.InfoF("Publish json: %s", string(body))
	msgBody := amqp.Publishing{
		Body:         body,
		DeliveryMode: amqp.Persistent,
		Headers:      map[string]interface{}{},
	}
	if err := amqptracer.Inject(sp, msgBody.Headers); err != nil {
		return err
	}

	return producer.channel.Publish(
		producer.exchange,
		routerKey,
		false,
		false,
		msgBody)
}

func (producer *Producer) Produce(routerKey string, data interface{}) error {
	body, err := json.Marshal(&data)
	if err != nil {
		return err
	}
	log.InfoF("Publish json: %s", string(body))
	msgBody := amqp.Publishing{
		Body:         body,
		DeliveryMode: amqp.Persistent,
		Headers:      map[string]interface{}{},
	}

	return producer.channel.Publish(
		producer.exchange,
		routerKey,
		false,
		false,
		msgBody)
}

func (producer *Producer) ProduceJsonWithoutLogs(ctx context.Context, routerKey string, data interface{}) error {
	sp := opentracing.SpanFromContext(ctx)
	defer sp.Finish()
	body, err := json.Marshal(&data)
	if err != nil {
		return err
	}
	msgBody := amqp.Publishing{
		Body:         body,
		DeliveryMode: amqp.Persistent,
		Headers:      map[string]interface{}{},
	}
	if err := amqptracer.Inject(sp, msgBody.Headers); err != nil {
		return err
	}
	return producer.channel.Publish(
		producer.exchange,
		routerKey,
		false,
		false,
		msgBody)
}

func (producer *Producer) ProduceProto(ctx context.Context, routerKey string, data proto.Message) error {
	sp := opentracing.SpanFromContext(ctx)
	defer sp.Finish()
	body, err := proto.Marshal(data)
	if err != nil {
		return err
	}
	msgBody := amqp.Publishing{
		Body:         body,
		DeliveryMode: amqp.Persistent,
		Headers:      map[string]interface{}{},
	}
	if err := amqptracer.Inject(sp, msgBody.Headers); err != nil {
		return err
	}
	return producer.channel.Publish(
		producer.exchange,
		routerKey,
		false,
		false,
		msgBody)
}

func (producer *Producer) Close() error {
	return producer.channel.Close()
}
