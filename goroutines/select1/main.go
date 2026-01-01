package main

import (
	"fmt"
	"time"
)

func operation1(channel chan string) {
	time.Sleep(1 * time.Second)
	channel <- "Operation 1"
}
func operation2(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Operation 2"
}
func main() {
	opt1 := make(chan string)
	opt2 := make(chan string)

	go operation1(opt1)
	go operation2(opt2)

	select {
	case choice1 := <-opt1:
		fmt.Println(choice1)
	case choice2 := <-opt2:
		fmt.Println(choice2)
	}
}
