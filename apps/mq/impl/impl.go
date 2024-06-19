package impl

import (
	"encoding/json"
	"sync"

	mqconfig "github.com/IanZC0der/go-myblog/apps/mq"
	"github.com/IanZC0der/go-myblog/ioc"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
)

type MQClient struct {
	Connection     *amqp.Connection
	Channel        *amqp.Channel
	resultChannels map[string]chan interface{}
	lock           sync.Mutex
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
		Connection:     conn,
		Channel:        ch,
		resultChannels: map[string]chan interface{}{},
	}, nil
}

func (mq *MQClient) Close() {
	mq.Channel.Close()
	mq.Connection.Close()

}

func GetMQClient() *MQClient {
	return ioc.DefaultControllerContainer().Get(mqconfig.AppName).(*MQClient)
}

func init() {
	ioc.DefaultControllerContainer().Register(&MQClient{
		resultChannels: map[string]chan interface{}{},
	})
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
	mq.resultChannels = make(map[string]chan interface{})

	_, err = mq.Channel.QueueDeclare(
		mqconfig.CREATE_BLOG_QUEUE,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	return nil
}

func (mq *MQClient) Name() string {
	return mqconfig.AppName
}

func (mq *MQClient) Publish(c *gin.Context, queueName string, body interface{}, resultChan chan interface{}) error {

	data, err := json.Marshal(body)
	if err != nil {
		return err
	}

	err = mq.Channel.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		},
	)
	if err != nil {
		return err
	}

	mq.StoreResultChannel(queueName, resultChan)
	return nil
}

func (mq *MQClient) Consumer(queueName string) (<-chan amqp.Delivery, error) {
	return mq.Channel.Consume(
		queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
}

func (mq *MQClient) StoreResultChannel(queueName string, resultChan chan interface{}) {
	mq.lock.Lock()
	defer mq.lock.Unlock()
	if mq.resultChannels == nil {
		mq.resultChannels = make(map[string]chan interface{})
	}
	mq.resultChannels[queueName] = resultChan
}

func (mq *MQClient) RetrieveResultChannel(queueName string) chan interface{} {
	mq.lock.Lock()
	defer mq.lock.Unlock()
	if mq.resultChannels == nil {
		return nil
	}
	resultChan, ok := mq.resultChannels[queueName]
	if !ok {
		return nil
	}
	delete(mq.resultChannels, queueName)
	return resultChan
}
