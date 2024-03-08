package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Number of CPUs:", runtime.NumCPU())
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))

	memStats := &runtime.MemStats{}
	runtime.ReadMemStats(memStats)
	fmt.Printf("Total memory: %v bytes\n", memStats.TotalAlloc)
	fmt.Printf("Memory allocated to heap objects: %v bytes\n", memStats.HeapAlloc)

	fmt.Println("Go version:", runtime.Version())
	fmt.Println("GOROOT:", runtime.GOROOT())
	fmt.Println("GOOS:", runtime.GOOS)
	fmt.Println("GOARCH:", runtime.GOARCH)
}