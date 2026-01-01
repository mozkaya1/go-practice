package main

import (
	"fmt"
	"sync"
	"time"
)

func displayContactList(wg *sync.WaitGroup) {
	contactList := []int{9011, 9022, 9033, 9044, 9055}
	for _, element := range contactList {
		time.Sleep(1 * time.Second) //causes delay of 1 second at each iteration
		fmt.Printf("%d ", element)
	}
	wg.Done() //Line3: WaitGroup counter decrements by 1
}
func displayContactNames(wg *sync.WaitGroup) {
	contactNames := []string{"John", "Tom", "Cindy", "Valerie", "Dexter"}
	for _, element := range contactNames {
		time.Sleep(1 * time.Second) //causes delay of 1 second at each iteration
		fmt.Printf("%s ", element)
	}
	wg.Done() //Line4: WaitGroup counter decrements by 1
}
func main() {
	var wg sync.WaitGroup       //Line1: Creating Instance of WaitGroup using "sync.waitGroup"
	wg.Add(2)                   //Line2: WaitGroup counter is set to 2
	go displayContactList(&wg)  //Function to print phone numbers
	go displayContactNames(&wg) //Function to print first contact names
	wg.Wait()                   //Line5: Wait() blocks the main goroutine until WaitGroup counter reaches 0
	fmt.Println("End-of-Main")
}
