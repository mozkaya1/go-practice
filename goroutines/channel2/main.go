package main

import (
	"fmt"
)

func main() {
	number := 3
	mainChannel := make(chan int)
	go square(number, mainChannel)
	valueReturned, status := <-mainChannel //Line1: Receiving value from recently closed mainChannel
	fmt.Println(valueReturned, status)     //status==true signifies channel is open
	go cube(number, mainChannel)
	valueReturned, status = <-mainChannel //Line2: Receiving zero value from already closed mainChannel
	fmt.Println(valueReturned, status)    //status==false signifies channel is closed
}
func square(num int, ch chan int) {
	ch <- num * num
	close(ch) //Line3: Closing mainChannel usign close()
}
func cube(num int, ch chan int) {
	fmt.Println(num * num * num)
}
