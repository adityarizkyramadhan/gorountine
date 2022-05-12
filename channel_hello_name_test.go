package gorountine

import (
	"fmt"
	"testing"
	"time"
)

func TestHelloChannel(t *testing.T) {
	var channel = make(chan string, 1)
	channel <- "Aditya"
	fmt.Println("Proses pemindahan data")
	fmt.Print("Loading")
	go helloNamaChannel(channel)
	for i := 0; i < 3; i++ {
		fmt.Print(".")
		time.Sleep(1 * time.Second)
	}
	data := <-channel
	fmt.Println("\nProses selesai")
	fmt.Println("Data :", data)
	close(channel)
}

func helloNamaChannel(channel chan string) {
	data := <-channel
	channel <- fmt.Sprintf("Hello %s", data)
}
