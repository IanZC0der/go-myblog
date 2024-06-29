package impl_test

import (
	"os"
	"strconv"
	"testing"

	// "time"

	mqconfig "github.com/IanZC0der/go-myblog/apps/mq"
	"github.com/IanZC0der/go-myblog/apps/mq/impl"
	"github.com/IanZC0der/go-myblog/ioc"
	"github.com/IanZC0der/go-myblog/test"

	"github.com/joho/godotenv"
	// "github.com/caarlos0/env/v6"
	// "os"
	// "github.com/caarlos0/env/v6"
)

var client *impl.MQClient

func init() {
	test.DevelopmentSetup()
	client = ioc.DefaultControllerContainer().Get(mqconfig.AppName).(*impl.MQClient)
}

func TestMQConfigFromENV(t *testing.T) {
	// t.Log(client)
	// aConfig := mqconfig.DefaultMQConfig()
	// if err := env.Parse(aConfig.RabbitMQ); err != nil {
	// 	t.Fatal(err)
	// }
	// t.Log(aConfig.RabbitMQ.Host)
	// t.Log(aConfig.RabbitMQ.Port)
	// t.Log(aConfig.RabbitMQ.User)
	// t.Log(aConfig.RabbitMQ.Password)
	// t.Log(os.Getenv("RABBITMQ_HOST"))
	err := godotenv.Load("../../../.env")
	if err != nil {
		t.Fatalf("Error loading .env file")
	}
	rabbitmqHost := os.Getenv("RABBITMQ_HOST")
	rabbitmqPort := os.Getenv("RABBITMQ_PORT")
	rabbitmqUser := os.Getenv("RABBITMQ_USER")
	rabbitmqPassword := os.Getenv("RABBITMQ_PASSWORD")
	t.Log(rabbitmqHost)
	t.Log(rabbitmqPort)
	t.Log(rabbitmqUser)
	t.Log(rabbitmqPassword)
	aConfig := mqconfig.DefaultMQConfig()

	aConfig.RabbitMQ.Host = rabbitmqHost
	aConfig.RabbitMQ.Port, _ = strconv.Atoi(rabbitmqPort)
	aConfig.RabbitMQ.User = rabbitmqUser
	aConfig.RabbitMQ.Password = rabbitmqPassword
	t.Log(aConfig.RabbitMQ.Host)
	t.Log(aConfig.RabbitMQ.Port)
	t.Log(aConfig.RabbitMQ.User)
	t.Log(aConfig.RabbitMQ.Password)

}

// func TestMQClient(t *testing.T) {

// 	defer client.Close()

// 	queueName := "testQueue"
// 	testMessage := "Hello, RabbitMQ!"

// 	// Assuming client is set up to declare queues and publish/consume as before
// 	_, err := client.Channel.QueueDeclare(queueName, true, false, false, false, nil)
// 	if err != nil {
// 		t.Fatalf("Failed to declare a queue: %v", err)
// 	}
// 	if err := client.Publish(queueName, []byte(testMessage)); err != nil {
// 		t.Fatalf("Failed to publish a message: %v", err)
// 	}

// 	msgs, err := client.Channel.Consume(queueName, "", true, true, false, false, nil)
// 	if err != nil {
// 		t.Fatalf("Failed to consume the messages: %v", err)
// 	}

// 	timeout := time.After(5 * time.Second)
// 	select {
// 	case d := <-msgs:
// 		if string(d.Body) != testMessage {
// 			t.Errorf("Got wrong message: %s", d.Body)
// 		}
// 	case <-timeout:
// 		t.Error("Test timed out waiting for the message")
// 	}
// }
