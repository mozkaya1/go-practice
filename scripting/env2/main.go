package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Get all environment variables
	envVars := os.Environ()
	// Print each environment variable
	for _, envVar := range envVars {
		pair := strings.SplitN(envVar, "=", 2)
		fmt.Printf("%s: %s\n", pair[0], pair[1])
	}
}
