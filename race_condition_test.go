package gorountine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	x := 0
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				x++
			}
		}()
	}
	time.Sleep(2 * time.Second)
	fmt.Println("Counter :", x)
}

func TestRaceConditionWithMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			mutex.Lock()
			for j := 0; j < 100; j++ {
				x++
			}
			mutex.Unlock()
		}()
	}
	time.Sleep(2 * time.Second)
	fmt.Println("Counter :", x)
}
