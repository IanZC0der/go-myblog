package mqconfig

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

var mqconfig = DefaultMQConfig()

type RabbitMQ struct {
	Host     string `json:"host" toml:"host" env:"RABBITMQ_HOST"`
	Port     int    `json:"port" toml:"port" env:"RABBITMQ_PORT"`
	User     string `json:"user" toml:"user" env:"RABBITMQ_USER"`
	Password string `json:"password" toml:"password" env:"RABBITMQ_PASSWORD"`
}

type MQConfig struct {
	RabbitMQ *RabbitMQ `json:"rabbitmq" toml:"rabbitmq"`
}

func DefaultMQConfig() *MQConfig {
	return &MQConfig{
		RabbitMQ: &RabbitMQ{
			Host:     "localhost",
			Port:     5672,
			User:     "guest",
			Password: "guest",
		},
	}
}

func LoadMQConfigFromENV() error {
	return env.Parse(mqconfig)
}

func (mq *RabbitMQ) GetAddress() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d",
		mq.User,
		mq.Password,
		mq.Host,
		mq.Port,
	)
}
