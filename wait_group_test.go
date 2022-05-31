package golanggoroutine

import (
	"sync"
	"testing"
	"time"
)

func RunAsync(group *sync.WaitGroup,number int) {
	defer group.Done()
	group.Add(1)
	
	time.Sleep(1 * time.Second)
	println("Perulangan ke-",number)
}

func TestWaitGroup(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		go RunAsync(&group,i)
	}

	group.Wait()
	println("Wait Group Done")
}
