package golanggoroutine

import (
	"sync"
	"testing"
)

func TestPool(t *testing.T) {
	var pool = sync.Pool{
		New: func() interface{} {
			return "New Pool"
		},
	}
	group := sync.WaitGroup{}

	pool.Put("Apel")
	pool.Put("Mangga")
	pool.Put("Jeruk")

	for i := 0; i < 10; i++ {
		go func() {
			group.Add(1)

			data := pool.Get().(string)
			println(data)
			// pool.Put(data)

			group.Done()
		}()
	}

	group.Wait()
	println("Pool Done")
}
