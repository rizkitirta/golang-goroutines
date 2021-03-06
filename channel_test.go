package golanggoroutine

import (
	"fmt"
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

// buffer channel atau limit memasukan data ke channel
func TestBufferChannel(t *testing.T) {
	channel := make(chan string, 2) // 2 adalah limit 
	defer close(channel)

	go func() {
		channel <- "data 1"
		channel <- "data 2"
	}()
	go func() {
		// data diambil sesuai urutan saat memasukan data ke channel
		data1 := <-channel
		data2 := <-channel
		println(data1)
		println(data2)
	}()

	time.Sleep(2 * time.Second)
	println("channel selesai")
}


func TestRangeChannel(t *testing.T)  {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- fmt.Sprintf("perulangan ke %d", i)
		}
		close(channel)
	}()

	for data := range channel {
		println("Menerima data " + data)
	}
}

func TestSelectChannel(t *testing.T)  {
	channel_1 := make(chan string)
	channel_2 := make(chan string)

	go GiveMeResponse(channel_1)
	go GiveMeResponse(channel_2)

	counter := 0
	for {
		select {
			case data := <-channel_1:
				println("Menerima data dari channel_1 " + data)
				counter++
			case data := <-channel_2:
				println("Menerima data dari channel_2 " + data)
				counter++
			default:
				println("Menunggu data")
		}

		if counter == 2 {break};
	}

	
}
