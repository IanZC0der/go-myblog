package impl_test

import (
	// "testing"
	// "time"

	mqconfig "github.com/IanZC0der/go-myblog/apps/mq"
	"github.com/IanZC0der/go-myblog/apps/mq/impl"
	"github.com/IanZC0der/go-myblog/ioc"
	"github.com/IanZC0der/go-myblog/test"
)

var client *impl.MQClient

func init() {
	test.DevelopmentSetup()
	client = ioc.DefaultControllerContainer().Get(mqconfig.AppName).(*impl.MQClient)
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
