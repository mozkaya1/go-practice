package main

import (
	"fmt"
	"sync"
)

type DigiWalletAccount struct {
	phoneNo int64
	balance int
	mutex   *sync.Mutex //Mutex pointer struct field
}

func (wallet *DigiWalletAccount) depositAccount(amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer wallet.mutex.Unlock()
	wallet.mutex.Lock()
	wallet.balance += amount

}

func (wallet *DigiWalletAccount) withDraw(amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer wallet.mutex.Unlock()

	wallet.mutex.Lock()
	wallet.balance -= amount
}

func main() {
	mutex := &sync.Mutex{}
	myAccount := DigiWalletAccount{555555, 10000, mutex}

	var wg sync.WaitGroup

	for range 1000 {
		wg.Add(2)
		go myAccount.depositAccount(100, &wg)
		go myAccount.withDraw(50, &wg)
	}
	wg.Wait()
	fmt.Printf("final balance of the wallet:%v\n", myAccount.balance)
}
