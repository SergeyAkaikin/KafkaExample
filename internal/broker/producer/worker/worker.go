package worker

import (
	"encoding/json"
	"log/slog"
	"math/rand"
	"time"

	"task/internal/message"
)

func Generator(out chan<- []byte, seconds int) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		num := rand.Intn(200)
		msg := message.New(num)
		value, err := json.Marshal(&msg)
		if err != nil {
			slog.Info("Marshaling problem", "error", err)
		}

		time.Sleep(time.Second * time.Duration(seconds))
		out <- value

	}
}
