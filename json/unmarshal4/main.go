package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
}

type Person struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Address Address `json:"address"`
}

func main() {
	data := []byte(`{"name":"Bob","age":35,"address":{"street":"123 Main St","city":"Springfield"}}`)
	var person Person
	err := json.Unmarshal(data, &person)
	if err != nil {
		panic(err)
	}
	fmt.Println(person)
	fmt.Println(person.Name, person.Age, person.Address.Street, person.Address.City)
}
