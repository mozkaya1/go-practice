package main

import (
	"encoding/json"
	"fmt"
)

type customer struct {
	Name     string
	Email    string
	Contacts []int
	Address  map[string]string
}

func main() {
	var customer1 customer
	jsonString := `{"Name":"John","Email":"john117@bungie.com","Contacts":[8989898990,8989898991,8989898992],"Address":{"city":"Colombus","country":"US","house":"12-B, Nixon Residency","state":"Ohio","street":"Sharlton Avenue"}}`
	err := json.Unmarshal([]byte(jsonString), &customer1) // unmarshal for decoding json as string
	if err != nil {
		panic(err)
	}
	fmt.Println(customer1)
}
