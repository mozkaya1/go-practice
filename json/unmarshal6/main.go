package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type CustomType interface {
	GetName() string
	GetAge() int
}

type CustomTypeImpl struct {
	person Person
}

func (c *CustomTypeImpl) GetName() string {
	return c.person.Name
}

func (c *CustomTypeImpl) GetAge() int {
	return c.person.Age
}

func main() {
	jsonData := []byte(`{"name": "John", "age": 30}`)
	var p CustomTypeImpl
	err := json.Unmarshal(jsonData, &p.person)
	if err != nil {
		panic(err)
	}
	var i CustomType = &p
	fmt.Println(i.GetName(), i.GetAge())
}
