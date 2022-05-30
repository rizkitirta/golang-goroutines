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

func TestChannelAsParam(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	// Menerima/menunggu data dari channel
	data := <-channel
	println(data)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(1 * time.Second)
	// mengirim data ke channel
	channel <- "Channel as parameter"
}
