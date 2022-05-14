package gorountine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type identitas struct {
	nama string
}

func TestPool(t *testing.T) {
	pool := &sync.Pool{}
	wg := &sync.WaitGroup{}
	nama := [...]identitas{{nama: "Aditya"}, {nama: "Rizky"}, {nama: "Ramadhan"}}
	pool.Put(&nama[0])
	pool.Put(&nama[1])
	pool.Put(&nama[2])
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			data := pool.Get()
			time.Sleep(time.Microsecond)
			if data != nil {
				fmt.Println(data.(*identitas).nama)
			}
			pool.Put(data)
			fmt.Println("Done")
			wg.Done()
		}()
		time.Sleep(2 * time.Second)
	}
	time.Sleep(10 * time.Second)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			data := pool.Get()
			time.Sleep(time.Microsecond)
			if data != nil {
				fmt.Println(data.(*identitas).nama)
			}
			pool.Put(data)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Selesai")
}
