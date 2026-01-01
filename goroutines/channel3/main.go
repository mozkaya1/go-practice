package main

import (
	"fmt"
)

func main() {
	number := 3
	mainChannel := make(chan int)
	go cube(number, mainChannel)
	valueReturned, status := <-mainChannel
	fmt.Println(valueReturned, status)
}
func cube(num int, ch chan int) {
	ch <- num * num * num
	close(ch) //Closing channel for first time
	close(ch) //Closing an already closed channel
}
