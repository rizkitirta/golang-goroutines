package golanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func SayHello(name string) {
	for i := 0; i < 100; i++ {
		println(i)
	}
	t := "Hello " + name
	println(t)
}

func TestGoroutine(t *testing.T) {
	go SayHello("Tirta")
	fmt.Println("Test Goroutine")

	time.Sleep(1 * time.Second)
}
