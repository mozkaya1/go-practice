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
	mainChannel <- 5 //Sending values to closed mainChannel
}
func cube(num int, ch chan int) {
	ch <- num * num * num
	close(ch) //mainChannel gets closed here
}
