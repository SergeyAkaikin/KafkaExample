package test

import (
	"encoding/json"
	"log"
	"math/rand"
	"testing"
	"time"

	"task/internal/broker/consumer"
	"task/internal/broker/producer"
	"task/internal/message"
)

var kafkaHost = `kafka:29092`
var topic = `test`
var group = `test`

func TestProduceConsume(t *testing.T) {

	p, err := producer.NewKafka(kafkaHost, topic)

	if err != nil {
		t.Errorf("Problem with connection to broker, %v", err)
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	value1 := message.New(r.Intn(200))

	msg, err := json.Marshal(&value1)

	if err != nil {
		t.Errorf("Problem with marshaling value1, %v", err)
	}

	if err = p.ProduceMsg(msg); err != nil {
		t.Errorf("Problem with producing, %v", err)
	}

	c, err := consumer.NewKafka(group, kafkaHost, []string{topic})
	if err != nil {
		t.Errorf("Problem with consumer connection to broker, %v", err)
	}

	if err = c.Subscribe(); err != nil {
		t.Errorf("Problem with consumer subscribing, %v", err)
	}

	for {
		value2, err := c.Consume()

		if err == nil {

			log.Println(value1, value2)

			if value1 != value2 {
				t.Errorf("value1 != value2, value1=%v, value2=%v", value1, value2)
			}
			break
		}

	}

}
