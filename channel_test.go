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

// hanya dapat mengirim data ke channel
func OnlyIn(channel chan<- string) {
	time.Sleep(1 * time.Second)
	channel <- "hanya dapat mengirim data ke channel"
}

// Hanya dapat Menerima/menunggu data dari channel
func OnlyOut(channel <-chan string) {
	data := <-channel
	println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(2 * time.Second)
	println("channel selesai")
}
