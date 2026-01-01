package main

import (
	"fmt"
	"sync"
	"time"
)

// Global variables for demonstration
// rwMutex protects access to the count variable
// Multiple readers can access count simultaneously, but only one writer at a time
var (
	rwMutex sync.RWMutex
	count   int
)

// read demonstrates a reader function that acquires a read lock
// Multiple readers can execute this function concurrently
func read() {
	rwMutex.RLock()                      // Acquire read lock
	defer rwMutex.RUnlock()              // Release read lock when function returns
	fmt.Println("Reading value:", count) // Read the shared count variable
	time.Sleep(1 * time.Second)          // Simulate some work being done
}

// write demonstrates a writer function that acquires a write lock
// Only one writer can execute this function at a time
// Writers block all other readers and writers
func write() {
	rwMutex.Lock()         // Acquire write lock
	defer rwMutex.Unlock() // Release write lock when function returns
	count++                // Modify the shared count variable
	fmt.Println("Writing value:", count)
	time.Sleep(1 * time.Second) // Simulate some work being done
}
func main() {
	// Launch 3 pairs of reader and writer goroutines
	for i := 0; i < 3; i++ {
		go read()  // Start a reader goroutine
		go write() // Start a writer goroutine
	}
	// Wait for goroutines to complete
	// In a real application, you would use WaitGroup instead of Sleep
	time.Sleep(5 * time.Second)
}
