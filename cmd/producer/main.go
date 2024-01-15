package main

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"task/internal/broker"
	"task/internal/broker/producer"
	"task/internal/broker/producer/worker"
	"task/internal/config"
)

func main() {

	conf := config.Body{}
	var loader config.Loader = config.JsonLoader{}

	if err := loader.Load(config.Path, &conf); err != nil {
		log.Fatalln(err)
	}

	closeSignal := make(chan os.Signal, 1)
	signal.Notify(closeSignal, syscall.SIGINT, syscall.SIGTERM)

	waitSignal := make(chan struct{})

	producers := make([]broker.Producer, 0, conf.ProducerNum)

	for i := 0; i < conf.ProducerNum; i++ {
		newProducer, err := producer.NewKafka(conf.Kafka, conf.Topic)
		if err == nil {
			producers = append(producers, newProducer)
			defer newProducer.Close()
		} else {
			slog.Warn("Problem with creating producer", "error", err)
		}
	}

	for _, p := range producers {
		go p.ListenEvents()

		msg := make(chan []byte)
		defer close(msg)

		go worker.Generator(msg, conf.Time)

		go p.StartProduce(msg, waitSignal)

		defer p.Flush(15 * 1000)
	}

	for {
		select {
		case <-closeSignal:
			for i := 0; i < len(producers); i++ {
				waitSignal <- struct{}{}
			}
			return
		case <-time.After(time.Second * 1):
		}
	}

}
