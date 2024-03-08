package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("\nCPU Information:")
	fmt.Println("Number of CPUs:", runtime.NumCPU())
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))

	fmt.Println("\nGO Information:")
	fmt.Println("Go version:", runtime.Version())
	fmt.Println("GOROOT:", runtime.GOROOT())
	fmt.Println("GOOS:", runtime.GOOS)
	fmt.Println("GOARCH:", runtime.GOARCH)
}