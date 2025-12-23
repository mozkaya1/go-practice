package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Name string
	Id   int
	Role int
}

type Marshaler interface {
	MarshalJSON() ([]byte, error)
}

func (e Employee) MarshalJSON() ([]byte, error) {
	type Alias Employee // create a type alias for the struct to avoid infinite recursion
	RoleMapping := map[int]string{2: "Operations Executive", 3: "Systems Engineer", 4: "Technology Analyst", 5: "Technology Lead"}
	return json.Marshal(&struct {
		Name string `json:"employeeName"`
		Id   int    `json:"employeeId"`
		Role string `json:"jobRole"`
	}{
		Name: e.Name,
		Id:   e.Id,
		Role: RoleMapping[e.Role],
	})
}

func main() {
	emp1 := Employee{Name: "Ramon", Id: 13456, Role: 3}
	jsonEncode, err := json.Marshal(emp1) // marshal for converting to json format
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonEncode))
}
