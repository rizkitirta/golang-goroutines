package golanggoroutine

import (
	"sync"
	"testing"
	"time"
)


func RunAsync(group *sync.WaitGroup, number int) {
	defer group.Done()
	group.Add(1)

	time.Sleep(1 * time.Second)
	println("Perulangan ke-", number)
}

// Sync WaitGroup digunakan untuk menunggu sampai semua goroutine selesai
func TestWaitGroup(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		go RunAsync(&group, i)
	}

	group.Wait()
	println("Wait Group Done")
}


// sync.Once digunakan untuk meng-execute sebuah fungsi hanya sekali
var counter int = 0
func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)
			once.Do(OnlyOnce)
			group.Done()
		}()
	}

	group.Wait()
	println("Counter:", counter)
}
