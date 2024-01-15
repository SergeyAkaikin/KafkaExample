package main

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"time"

	"task/internal/broker"
	"task/internal/broker/consumer"
	"task/internal/config"
)

func main() {

	var loader config.Loader = config.JsonLoader{}

	conf := config.Body{}

	if err := loader.Load(config.Path, &conf); err != nil {
		log.Fatalln(err)
	}

	closeSignal := make(chan os.Signal, 1)
	signal.Notify(closeSignal, syscall.SIGINT, syscall.SIGTERM)

	waitSignal := make(chan struct{})

	consumers := make([]broker.Consumer, 0, conf.ConsumerNum)

	for i := 0; i < conf.ConsumerNum; i++ {
		groupId := conf.GroupId + strconv.Itoa(i)
		newConsumer, err := consumer.NewKafka(groupId, conf.Kafka, []string{conf.Topic})
		if err == nil {
			consumers = append(consumers, newConsumer)
			defer newConsumer.Close()
		} else {
			slog.Warn("Problem with creating consumer", "error", err)
		}
	}

	for _, c := range consumers {
		go c.StartConsume(waitSignal, conf.MessagesCount)
	}

	for {
		select {
		case <-closeSignal:
			for i := 0; i < len(consumers); i++ {
				waitSignal <- struct{}{}
			}
			return
		case <-time.After(time.Second * 1):
		}
	}
}
