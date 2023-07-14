package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type RabbitMQHook struct {
	ConnectionString string
	Exchange         string
	RoutingKey       string
	LogLevel         logrus.Level
}

func CreateRabbitMQHook(connectionString, exchange, routingKey string, logLevel logrus.Level) *RabbitMQHook {
	return &RabbitMQHook{
		ConnectionString: connectionString,
		Exchange:         exchange,
		RoutingKey:       routingKey,
		LogLevel:         logLevel,
	}
}

func (hook *RabbitMQHook) Fire(entry *logrus.Entry) error {
	msg := entry.Message
	if err := hook.publishMessage(msg); err != nil {
		return err
	}
	return nil
}

func (hook *RabbitMQHook) Levels() []logrus.Level {
	return logrus.AllLevels[:hook.LogLevel+1]
}

func (hook *RabbitMQHook) publishMessage(msg string) error {
	conn, err := amqp.Dial(hook.ConnectionString)
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	err = ch.Publish(
		hook.Exchange,
		hook.RoutingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		},
	)
	if err != nil {
		return err
	}

	return nil
}

