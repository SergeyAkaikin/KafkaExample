package broker

type Producer interface {
	StartProduce(msg <-chan []byte, signal <-chan struct{})
	ListenEvents()
	Close()
	Flush(int) int
}

type Consumer interface {
	StartConsume(signal <-chan struct{}, msgCount int)
}
