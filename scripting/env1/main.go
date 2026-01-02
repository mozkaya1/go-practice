package main

import (
	"fmt"
	"os"
)

func main() {
	// Set a new environment variable
	err := os.Setenv("MY_VAR", "Hello, Golang!")
	if err != nil {
		fmt.Println("Error setting environment variable:", err)
		return
	}
	// Get the value of the new environment variable
	myVar := os.Getenv("MY_VAR")
	fmt.Println("MY_VAR:", myVar)
}
