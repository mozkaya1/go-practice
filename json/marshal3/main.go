package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Id          int       `json:"customerId"`
	Name        string    `json:"customerName"`
	PhoneNumber int       `json:"phoneNumber"`
	ContactList []contact `json:"contacts"`
}

type contact struct {
	Name    string `json:"contactName"`
	PhoneNo int    `json:"phoneNumber"`
}

func main() {
	Users := user{1001, "John", 8989898999, []contact{{"Jack", 9989898923}, {"Mary", 9989898453}, {"Tom", 9989898898}}}
	jsonObject, err := json.Marshal(Users)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonObject)) //To display byte array as a string
}
