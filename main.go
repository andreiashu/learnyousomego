package main

import (
	"fmt"
	"runtime"
)

func main() {
	nCPU := runtime.NumCPU()

	fmt.Println("Number of CPUs: ", nCPU)
}
