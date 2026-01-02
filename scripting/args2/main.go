package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	// Define the operation flag
	operation := flag.String("op", "sum", "Operation to perform: sum, sub, mul, div")
	flag.Parse()
	// Remaining arguments after flags
	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("Usage: go run main.go -op=<operation> <num1> <num2>")
		return
	}
	// Convert arguments to integers
	num1, err1 := strconv.Atoi(args[0])
	num2, err2 := strconv.Atoi(args[1])
	if err1 != nil || err2 != nil {
		fmt.Println("Please provide two valid integers.")
		return
	}
	// Perform the selected operation
	switch *operation {
	case "sum":
		fmt.Printf("Result: %d + %d = %d\n", num1, num2, num1+num2)
	case "sub":
		fmt.Printf("Result: %d - %d = %d\n", num1, num2, num1-num2)
	case "mul":
		fmt.Printf("Result: %d * %d = %d\n", num1, num2, num1*num2)
	case "div":
		if num2 == 0 {
			fmt.Println("Error: Division by zero is not allowed.")
		} else {
			fmt.Printf("Result: %d / %d = %d\n", num1, num2, num1/num2)
		}
	default:
		fmt.Println("Invalid operation. Use one of: sum, sub, mul, div")
	}
}
