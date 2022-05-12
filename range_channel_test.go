package gorountine

import (
	"fmt"
	"testing"
)

func TestLoopChannel(t *testing.T) {
	var channel = make(chan int, 1)
	go func() {
		for i := 0; i < 10000; i++ {
			channel <- i
		}
		close(channel)
	}()
	for i := range channel {
		fmt.Println(i)
	}
}
