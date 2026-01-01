package main

import (
	"context"
	"fmt"
	"time"
)

func fetchData(channel chan string) {
	start := time.Now()
	for i := 0; i < 5; i++ {
		time.Sleep(1500 * time.Millisecond)
		fmt.Printf("Processing request... (time lapsed: %v) \n", time.Since(start))
	}
	data := fmt.Sprintf("Data fetched after %v seconds", time.Since(start))
	channel <- data
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	channel := make(chan string)
	go fetchData(channel)
	time.Sleep(5 * time.Second)
	cancel()

	select {
	case data := <-channel:
		fmt.Println(data)
	case <-ctx.Done():
		fmt.Println("Fetch operation cancelled after 5 secs ....")

	}
}
