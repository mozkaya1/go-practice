package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
)

func allocateMemory() []*int {
	s := make([]*int, 0)
	for i := 0; i < 1e6; i++ {
		v := new(int)
		*v = i
		s = append(s, v)
	}
	return s
}
func main() {
	f, err := os.Create("mem.prof")
	if err != nil {
		fmt.Println("Could not create memory profile:", err)
		return
	}
	allocateMemory()
	runtime.GC() // Run garbage collection before profiling
	pprof.WriteHeapProfile(f)
	f.Close()
}

