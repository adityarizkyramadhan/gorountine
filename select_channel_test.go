package gorountine

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go giveMeNumber(ch1)
	go giveMeNumber(ch2)
	for {
		time.Sleep(time.Millisecond * 100)
		select {
		case v1 := <-ch1:
			fmt.Println("Data ch1 :", v1)
		case v2 := <-ch2:
			fmt.Println("Data ch3 :", v2)
		default:
			fmt.Println("Belum ada data")
			return
		}
	}
}

func giveMeNumber(num chan int) {
	num <- rand.Intn(100)
}
