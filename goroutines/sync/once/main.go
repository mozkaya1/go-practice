package main

import (
	"fmt"
	"sync"
	"time"
)

// once is used to ensure that initialization happens only once
// even across multiple goroutines
var (
	once sync.Once
)

// initialize simulates a one-time initialization process
// This could be database connection, config loading, etc.
func initialize() {
	fmt.Println("Initialization done")
}
func main() {
	// Launch 3 goroutines that will try to perform initialization
	for i := 0; i < 3; i++ {
		go func() {
			// once.Do() ensures the initialize() function is called exactly once
			// subsequent calls to Do() will not execute initialize() again
			once.Do(initialize)
		}()
	}
	// Wait for goroutines to complete
	// In a real application, you would use WaitGroup instead of Sleep
	time.Sleep(5 * time.Second)
}
