package main

import (
	"encoding/json"
	"fmt"
)

type Animal interface {
	Speak() string
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return "Meow!"
}

type Goat struct {
	Name string
}

func (g Goat) Speak() string {
	return "Maa!"
}

type Pet struct {
	Name   string
	Animal Animal
}

func main() {
	pet1 := Pet{Name: "Fido", Animal: Cat{Name: "Fido"}}
	jsonBytes, err := json.Marshal(pet1)
	if err != nil {
		panic(err)
	}
	jsonString := string(jsonBytes)
	fmt.Println(jsonString)

	pet1 = Pet{Name: "Fido", Animal: Cat{Name: "Fido"}}
	pet2 := Pet{Name: "Ramsey", Animal: Goat{Name: "Ramsey"}}
	pets := []Pet{pet1, pet2}
	jsonBytes, err = json.Marshal(pets)
	if err != nil {
		panic(err)
	}
	jsonString = string(jsonBytes)
	fmt.Println(jsonString)
	fmt.Println(pet1.Animal.Speak())
	fmt.Println(pet2.Animal.Speak())
}
