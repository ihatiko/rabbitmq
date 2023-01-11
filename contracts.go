package rabbitmq

import (
	"context"
	"github.com/streadway/amqp"
	"google.golang.org/protobuf/proto"
)

type IProducer interface {
	ProduceJson(ctx context.Context, routerKey string, data interface{}) error
	ProduceJsonWithoutLogs(ctx context.Context, routerKey string, data interface{}) error
	ProduceProto(ctx context.Context, routerKey string, data proto.Message) error
	Close() error
}

type IConsumer interface {
	Consume(handler func(context context.Context, d amqp.Delivery) bool) (*Channel, error)
}
