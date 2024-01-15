package config

import "os"

func init() {

	path, exists := os.LookupEnv("CONFIG")
	if exists {
		Path = path
	}
}

type Loader interface {
	Load(string, interface{}) error
}

var Path string

type Body struct {
	MessagesCount int    `json:"messagesCount"`
	Time          int    `json:"time"`
	ProducerNum   int    `json:"producerNum"`
	ConsumerNum   int    `json:"consumerNum"`
	Kafka         string `json:"kafka"`
	Topic         string `json:"topic"`
	GroupId       string `json:"groupId"`
}
