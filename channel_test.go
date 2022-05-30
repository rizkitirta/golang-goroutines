package golanggoroutine

import (
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		// sleep digunakan untuk simulasi saat data diproses
		time.Sleep(1 * time.Second)
		channel <- "Hello Tirta"

		println("Selesai mengirim data ke channel")
	}()

	data := <-channel
	println(data)
}
