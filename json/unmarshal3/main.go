package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var customer interface{}
	jsonInterface := `{"name":"John","age":80}`
	error := json.Unmarshal([]byte(jsonInterface), &customer) // unmarshal for decoding
	if error != nil {
		panic(error)
	}
	fmt.Println(customer)
	// Access fields of the unmarshalled data
	m := customer.(map[string]interface{}) //Using type assertion to store the decoded data into a map

	fmt.Println("Name:", m["name"].(string))
	fmt.Println("Age:", m["age"].(float64))
}
