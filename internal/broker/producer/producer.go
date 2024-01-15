package producer

import (
	"encoding/json"
	"log/slog"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaProducer struct {
	*kafka.Producer
	topic string
}

func NewKafka(broker string, topic string) (KafkaProducer, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
	return KafkaProducer{producer, topic}, err
}

func (p KafkaProducer) StartProduce(in <-chan []byte, signal <-chan struct{}) {

	for {
		select {
		case value := <-in:
			if err := p.ProduceMsg(value); err != nil {
				slog.Warn("Error", "", err)
			}
		case <-signal:
			return
		}
	}
}

type stats struct {
	MsgCnt        int `json:"txmsgs"`
	MsgSize       int `json:"txmsg_bytes"`
	RequestNumber int `json:"tx"`
}

func (p KafkaProducer) ListenEvents() {
	for e := range p.Events() {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				slog.Info("Delivery failed", "error", ev.TopicPartition)
			} else {
				slog.Info("Delivered message", "to", ev.TopicPartition)
			}
		case *kafka.Stats:
			stat := stats{}
			err := json.Unmarshal([]byte(ev.String()), &stat)
			if err != nil {
				slog.Warn("Unmarshalling stats problem", "ErrorInfo", err)
			} else {
				slog.Info("Entered", "stats", stat)
			}
		}
	}
}

func (p KafkaProducer) ProduceMsg(msg []byte) error {
	if err := p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &p.topic, Partition: kafka.PartitionAny},
		Value:          msg,
	}, nil); err != nil {
		return err
	}

	return nil
}
