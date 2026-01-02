package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Ensure exactly two numeric arguments are provided.
	fmt.Println(len(os.Args))
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <num1> <num2>")
		return
	}
	// Parse the first argument to an integer.
	num1, err1 := strconv.Atoi(os.Args[1])
	// Parse the second argument to an integer.
	num2, err2 := strconv.Atoi(os.Args[2])
	// If either conversion fails, notify the user and exit.
	if err1 != nil || err2 != nil {
		fmt.Println("Please provide two valid integers.")
		return
	}
	// Compute the sum of the two integers.
	sum := num1 + num2
	// Print the result to stdout.
	fmt.Printf("The sum of %d and %d is %d\n", num1, num2, sum)
}
