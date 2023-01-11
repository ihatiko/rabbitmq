package rabbitmq

import (
	"context"
	"fmt"
	"github.com/streadway/amqp"
	"testing"
	"time"
)

const (
	testConsumer1 = "testConsumer1"
	testConsumer2 = "testConsumer2"
	testProducer  = "testProducer"
	testRouterKey = "testRouterKey"
	testExchange  = "testExchange"
	appName       = "test"
)

type TestStruct struct {
	Field string `json:"field"`
}

func TestConnection_NewConnection(t *testing.T) {
	conProducer, err := NewConnection(testProducer, &Config{
		RetryTTL:    30,
		UserName:    "guest",
		Concurrency: 1,
		Host:        "localhost",
		Password:    "guest",
		Port:        "5672",
	})
	if err != nil {
		panic(err)
	}
	producer, err := conProducer.GetProducer(testExchange)
	if err != nil {
		panic(err)
	}

	conConsumer1, err := NewConnection(testConsumer1, &Config{
		RetryTTL:    30,
		UserName:    "guest",
		Concurrency: 1,
		Host:        "localhost",
		Password:    "guest",
		Port:        "5672",
	})
	if err != nil {
		panic(err)
	}
	conConsumer2, err := NewConnection(testConsumer2, &Config{
		RetryTTL:    30,
		UserName:    "guest",
		Concurrency: 1,
		Host:        "localhost",
		Password:    "guest",
		Port:        "5672",
	})
	if err != nil {
		panic(err)
	}

	consumer1, err := conConsumer1.GetConsumer(testExchange, testConsumer1, testRouterKey)
	if err != nil {
		panic(err)
	}

	consumer2, err := conConsumer2.GetConsumer(testExchange, testConsumer2, testRouterKey)
	if err != nil {
		panic(err)
	}

	consumer1.Consume(func(context context.Context, d amqp.Delivery) bool {
		fmt.Println(d.Headers)
		fmt.Println(string(d.Body), "consumer1")
		return false
	})

	consumer2.Consume(func(context context.Context, d amqp.Delivery) bool {
		fmt.Println(d.Headers)
		fmt.Println(string(d.Body), "consumer2")
		return false
	})
	err = producer.Produce(testRouterKey, &TestStruct{Field: "test"})
	if err != nil {
		panic(err)
	}
	for range time.Tick(time.Second * 1) {

	}
}
