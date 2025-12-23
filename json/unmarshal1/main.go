package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var contactsSlice []int
	jsonSlice := `[8989898990, 8989898991, 8989898992]`
	fmt.Printf("Before -- Type:%T\n", jsonSlice)
	error := json.Unmarshal([]byte(jsonSlice), &contactsSlice) // unmarshal for decoding json as slice
	if error != nil {
		panic(error)
	}
	fmt.Printf("After -- Type:%T\n Value:%v\n", contactsSlice, contactsSlice)
	var loginMap map[string]int
	jsonMap := `{ "ID1":8989898990 ,  "ID2":8989898991,  "ID3":8989898992 }`
	err := json.Unmarshal([]byte(jsonMap), &loginMap) // unmarshal for decoding json as map
	if err != nil {
		panic(err)
	}
	fmt.Println(loginMap)
}
