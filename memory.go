package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("\nMemory Information:")
	memStats := &runtime.MemStats{}
	runtime.ReadMemStats(memStats)
	fmt.Printf("Total memory: %v bytes\n", memStats.TotalAlloc)
	fmt.Printf("Memory allocated to heap objects: %v bytes\n", memStats.HeapAlloc)
}