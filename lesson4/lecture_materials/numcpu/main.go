package main

import (
	"fmt"
	"runtime"
)

func main() {
	// NumCPU returns the number of logical
	// CPUs usable by the current process.
	fmt.Println(runtime.NumCPU())
}
