package impl

import (
	mqconfig "github.com/IanZC0der/go-myblog/apps/mq"
	"github.com/IanZC0der/go-myblog/ioc"
	amqp "github.com/rabbitmq/amqp091-go"
)

type MQClient struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func NewClient(dsn string) (*MQClient, error) {
	conn, err := amqp.Dial(dsn)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	return &MQClient{
		Connection: conn,
		Channel:    ch,
	}, nil
}

func (mq *MQClient) Close() {
	mq.Channel.Close()
	mq.Connection.Close()
}

func init() {
	ioc.DefaultControllerContainer().Register(&MQClient{})
}

func (mq *MQClient) Init() error {
	address := mqconfig.DefaultMQConfig().RabbitMQ.GetAddress()
	conn, err := amqp.Dial(address)
	if err != nil {
		return err
	}
	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return err
	}

	mq.Connection = conn
	mq.Channel = ch
	return nil
}

func (mq *MQClient) Name() string {
	return mqconfig.AppName
}

func (mq *MQClient) Publish(queueName string, body []byte) error {
	_, err := mq.Channel.QueueDeclare(
		queueName,
		true,  // Durable
		false, // Delete when unused
		false, // Exclusive
		false, // No-wait
		nil,   // Arguments
	)
	if err != nil {
		return err
	}

	return mq.Channel.Publish(
		"",        // Exchange
		queueName, // Routing key
		false,     // Mandatory
		false,     // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	)
}
