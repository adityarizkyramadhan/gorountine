package gorountine

import (
	"fmt"
	"testing"
)

func printAngka(num int) {
	fmt.Println(num)
}
func TestHelloWorld(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go printAngka(i)
	}
}
