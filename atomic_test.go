package gorountine

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

func TestAtomic(t *testing.T) {
	var x int32 = 0
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				atomic.AddInt32(&x, 1)
			}
		}()
	}
	time.Sleep(2 * time.Second)
	fmt.Println("Counter :", x)
}
