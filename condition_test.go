package gorountine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var wg = sync.WaitGroup{}
var cond = sync.NewCond(&sync.Mutex{})

func waitCondition(value int) {
	cond.L.Lock()
	cond.Wait()
	fmt.Println("Gorountine ke", value)
	cond.L.Unlock()
	wg.Done()
}

func TestCondition(t *testing.T) {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go waitCondition(i)
	}
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(2 * time.Second)
			cond.Signal()
		}
	}()
	time.Sleep(2 * time.Second)
	// for i := 0; i < 10; i++ {
	// 	go func() {
	// 		time.Sleep(2 * time.Second)
	// 		cond.Broadcast()
	// 	}()
	// }
	wg.Wait()
}
