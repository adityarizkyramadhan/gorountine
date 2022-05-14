package gorountine

import (
	"fmt"
	"sync"
	"testing"
)

func addToMap(value int, mapSync *sync.Map, wg *sync.WaitGroup) {
	defer wg.Done()
	mapSync.Store(value, value)
}

func TestMapSync(t *testing.T) {
	mapSync := &sync.Map{}
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go addToMap(i, mapSync, &wg)
	}
	wg.Wait()

	mapSync.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}
