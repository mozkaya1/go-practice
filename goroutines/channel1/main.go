package main

import (
	"fmt"
)

func main() {
	tester := 3
	// This is where we "make" the channel, which can be used
	// to move the `int` value(s)
	var outChannel chan int = make(chan int) //Line1
	// We still run this function as a goroutine
	// But we also provide the channel that we made in Line1
	go square(tester, outChannel)
	// Once any value is received from outChannel to main goroutine, we print it to the console and proceed
	fmt.Println(<-outChannel + 3) //Line2
}

// This function now accepts a channel as its second argument
func square(num int, ch chan int) {
	result := num * num
	// The result value is sent from square goroutine to outChannel
	ch <- result //Line3
}
