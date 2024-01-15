package consumer

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"

	"task/internal/message"
)

type stats struct {
	MsgConsumed int `json:"rxmsgs"`
	MsgSize     int `json:"rxmsg_bytes"`
}
type KafkaConsumer struct {
	topics []string
	*kafka.Consumer
}

func NewKafka(groupId string, host string, topics []string) (KafkaConsumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":      host,
		"group.id":               groupId,
		"auto.offset.reset":      "earliest",
		"statistics.interval.ms": 30000,
	})

	return KafkaConsumer{topics, c}, err
}

func (c KafkaConsumer) Subscribe() (err error) {
	err = c.SubscribeTopics(c.topics, nil)
	return
}

func (c KafkaConsumer) StartConsume(signal <-chan struct{}, msgCount int) {

	if err := c.Subscribe(); err != nil {
		slog.Warn("Subscription problem", "error", err)
		return
	}

	buff := make([]message.Sync, 0, msgCount)

	for {
		if len(buff) == msgCount {
			log.Println(buff)
			buff = buff[:0]
		}

		select {
		case <-signal:
			return
		default:

			ev := c.Poll(100)

			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				value := message.Sync{}
				if err := json.Unmarshal(e.Value, &value); err != nil {
					slog.Warn("StartConsumer: error of unmarshaling message", "error", err)
				}
				buff = append(buff, value)
			case kafka.Error:
				if e.IsFatal() {
					slog.Warn("Kafka error", "", e)
				} else {
					continue
				}
			case *kafka.Stats:
				stat := stats{}
				err := json.Unmarshal([]byte(ev.String()), &stat)
				if err != nil {
					slog.Warn("Unmarshalling stats problem", "error", err)
				} else {
					slog.Info("Entered stats", "", stat)
				}
			}
		}
	}

}

func (c KafkaConsumer) Consume() (message.Sync, error) {

	ev := c.Poll(100)

	value := message.Sync{}

	if ev == nil {
		return value, fmt.Errorf("c.Poll is nil")
	}

	switch e := ev.(type) {
	case *kafka.Message:
		if err := json.Unmarshal(e.Value, &value); err != nil {
			return value, fmt.Errorf("consume: error of unmarshaling message; %w", err)
		}
	case kafka.Error:
		return value, e
	}

	return value, nil
}
