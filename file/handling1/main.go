package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	_ "strconv"
	"strings"
)

type Customer struct {
	Name    string
	Address string
	Age     int
}

func main() {
	// Read the content of the file into a byte slice
	content, err := ioutil.ReadFile("customer.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Convert byte slice to string and split it into lines
	fmt.Println(string(content))
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")

	// Iterate over each line and split it using comma separator
	for _, line := range lines {
		fields := strings.Split(line, ",")

		// Extract name, address, and age fields from the line
		name := strings.TrimSpace(fields[0])
		address := strings.TrimSpace(fields[1])
		ageStr := strings.TrimSpace(fields[2])

		// Convert age field to integer
		age, _ := strconv.Atoi(ageStr)

		// Create a new customer with extracted data
		customer := Customer{
			Name:    name,
			Address: address,
			Age:     age,
		}

		// Print the customer data
		fmt.Printf("%v\n", customer)
	}
}
