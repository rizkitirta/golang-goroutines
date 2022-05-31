package golanggoroutine

import (
	"sync"
	"testing"
)

func AddToMap(data *sync.Map, value int, group *sync.WaitGroup) {
	defer group.Done()
	
	group.Add(1)
	data.Store(value, value)
}

func TestMap(t *testing.T) {
	dataMap := sync.Map{}
	group := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		go AddToMap(&dataMap, i, &group)
	}

	group.Wait()
	
	dataMap.Range(func(key, value interface{}) bool {
		println(key.(int), value.(int))
		return true
	})
	println("Map Done")
}
