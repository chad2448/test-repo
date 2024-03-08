package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("\nGO Information:")
	fmt.Println("Go version:", runtime.Version())
	fmt.Println("GOROOT:", runtime.GOROOT())
	fmt.Println("GOOS:", runtime.GOOS)
	fmt.Println("GOARCH:", runtime.GOARCH)
}