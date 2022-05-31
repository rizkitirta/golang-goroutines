package golanggoroutine

import (
	"sync"
	"testing"
	"time"
)

var mutex sync.Mutex
var cond = sync.NewCond(&mutex)
var group = &sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)

	cond.L.Lock()
	cond.Wait()

	println("value:", value)

	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	for i := 0; i < 100; i++ {
		go WaitCondition(i)
	}

	go func ()  {
		for i := 0; i < 100; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal()
		}
	}()

	// Broadcast mengakibatkan semua goroutine otomatis berjalan setelah lock 
	// go func ()  {
	// 	time.Sleep(3 * time.Second)
	// 	cond.Broadcast()
	// }()

	group.Wait()
	println("Cond Done")
}