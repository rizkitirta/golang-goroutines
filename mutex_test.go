package golanggoroutine

import (
	"sync"
	"testing"
	"time"
)

// Mutex digunakan untuk mengunci variabel saat diakses oleh goroutine
func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex
	
	for i := 0; i < 1000; i++ {
		go func()  {
			for j := 0; j < 100; j++ {
				mutex.Lock()
				x++
				mutex.Unlock()
			}	
		}()
	}

	time.Sleep(5 * time.Second)
	println(x)
}