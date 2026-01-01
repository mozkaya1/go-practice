package main

import (
	"context"
	"fmt"
	"time"
)

func fetchData(channel chan string) {
	time.Sleep(14 * time.Second)
	channel <- "Data fecthed after 14 sec...."
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	channel := make(chan string)

	go fetchData(channel)

	select {
	case data := <-channel:
		fmt.Println(data)
	case <-ctx.Done():
		fmt.Println("Fetch operation Timeout after 5 secs ...")
	}
}
