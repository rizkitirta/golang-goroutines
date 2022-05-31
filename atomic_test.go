package golanggoroutine

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var number int64 = 0
	var group = sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		go func() {
			group.Add(1)
			for j := 0; j < 100; j++ {
				atomic.AddInt64(&number, 1)
			}
			group.Done()
		}()
	}

	group.Wait()
	println(number)
}
