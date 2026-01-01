package main

import (
	"fmt"
	"time"
)

type customer struct {
	customerId   int
	customerName string
}

func customerData1(chan1 chan customer) {
	time.Sleep(4 * time.Second)
	chan1 <- customer{1001, "Jack"}
}

func customerData2(chan2 chan customer) {
	time.Sleep(6 * time.Second)
	chan2 <- customer{1002, "Lucy"}
}

func customerData3(chan3 chan customer) {
	time.Sleep(6 * time.Second)
	chan3 <- customer{1003, "Harry"}
}

func main() {
	opt1 := make(chan customer)
	opt2 := make(chan customer)
	opt3 := make(chan customer)

	go customerData1(opt1)
	go customerData2(opt2)
	go customerData3(opt3)

	select {
	case choice1 := <-opt1:
		fmt.Println(choice1)
	case choice2 := <-opt2:
		fmt.Println(choice2)
	case choice3 := <-opt3:
		fmt.Println(choice3)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout Occurred")
	}
}
