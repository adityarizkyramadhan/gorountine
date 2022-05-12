package gorountine

import (
	"fmt"
	"sync"
	"testing"
)

func TestSyncWait(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		go runAsycncronus(&wg)
	}
	wg.Wait()
	fmt.Println("Selesai")
}

func runAsycncronus(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		fmt.Println(i)
	}
}
