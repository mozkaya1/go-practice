package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Address struct {
	Street string `json:"street2"`
	City   string `json:"city"`
}

type Person struct {
	Name    string  `json:"name"`
	Age     int     `json:"age,omitempty"`
	Address Address `json:"address,omitempty"`
}

type Person2 struct {
	Name string
	Age  int
}

func marshal2() {

	people := []Person2{
		{Name: "John Smith", Age: 30},
		{Name: "Jane Doe", Age: 25},
		{Name: "Bob Johnson", Age: 40},
	}
	jsonBytes, err := json.Marshal(people)
	if err != nil {
		panic(err)
	}
	jsonString := string(jsonBytes)
	fmt.Println(jsonString)

	// Senario 2
	data := map[string]string{
		"name":  "John Smith",
		"email": "john@example.com",
		"phone": "555-1234",
	}
	jsonBytes2, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	jsonString2 := string(jsonBytes2)
	fmt.Println(jsonString2)
}

func marshal1() {
	// Scenario 1
	// p := Person{
	// Name: "John Smith",
	// Address: Address{
	//     Street:  "123 Main St",
	//     City:    "New York",
	// },

	// Scenario 2

	p := Person{
		Name: "John Smith",
		Address: Address{
			Street: "123 Main St",
			City:   "New York",
		},
	}

	jsonBytes, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	jsonString := string(jsonBytes)
	fmt.Println(jsonString)
}
func main() {
	contactSlice := []int{21312321, 231232143, 123123213}
	jsonSlice, err := json.Marshal(contactSlice)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Encoded slice", string(jsonSlice))
	fmt.Printf("type %v", jsonSlice)
	addressMap := map[string]string{"house": "12-B, Nixon Residency", "street": "Sharlton Avenue", "city": "Colombus", "state": "Ohio", "country": "US"}

	jsonMap, err := json.Marshal(addressMap)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("encoded Map", string(jsonMap))
	fmt.Println("==============================")
	marshal1()
	fmt.Println("==============================")
	marshal2()
}
