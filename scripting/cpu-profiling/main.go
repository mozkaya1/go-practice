package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func heavyComputation() {
	for i := 0; i < 1e7; i++ {
		_ = i * i
	}
}
func main() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Println("Could not create CPU profile:", err)
		return
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	heavyComputation()
	time.Sleep(2 * time.Second)
}
