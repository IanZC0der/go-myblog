# Integrate RabbitMQ with API handlers

## RabbitMQ Client structure

```
type MQClient struct {
	Connection     *amqp.Connection
	Channel        *amqp.Channel
	resultChannels map[string]chan interface{}
	lock           sync.Mutex
}

```

- resultChannels can be used to store the results. For example, after a new blog is created, we can put it into resultChannels, the api handler can retrieve the blog from the corresponding channels. Each queue should have a separate channel

- mutex lock is used since multiple threads can try to update the channels at the same time

## workflow of rabbitmq with create blog api

- in the init() function, all the queues should be declared

- in the api handler function, create the result channel and encode the request. call the publish function to publish the message, pass in the channel and message. at last, try to retrieve the created blog from the channel

- publish function will handle the logic of save the channel into the client

- in the consumer function, consume the message and create the blog and add the blog to the channel with go routine
