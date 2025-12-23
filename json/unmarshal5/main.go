package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var namesSlice []string
	jsonSlice := `["Rajesh", "Rocky", "Shelly"]`
	// write the code to decode the given slice
	err := json.Unmarshal([]byte(jsonSlice), &namesSlice)
	fmt.Printf("Type Jsonslice:%T\n Data:%v\n", jsonSlice, jsonSlice)
	fmt.Printf("Type namesSlice:%T\n Data:%v\n", namesSlice, namesSlice)

	// Activity 2: Uncomment the below code snippet
	var loginMap map[string]int
	jsonMap := `{ "Infy@123":8989898990 ,  "Infy@321":8989898991,  "Infy@456":8989898992 }`
	err = json.Unmarshal([]byte(jsonMap), &loginMap)
	if err != nil {
		panic(err)
	}
	fmt.Println(loginMap)
}
