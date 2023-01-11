package rabbitmq

const (
	deadlyExchangePrefix = "X_DEAD_LETTER_EXCHANGE"
	defaultExchangeType  = "topic"
	xDeadLetterExchange  = "x-dead-letter-exchange"
	xDeadMessageTTL      = "x-message-ttl"
	innerExchange        = "inner"
)
