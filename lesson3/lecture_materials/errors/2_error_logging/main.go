package main

import (
	"fmt"
	"strconv"
)

func RunJob1(item string) error {
	_, err := strconv.Atoi(item)
	if err != nil {
		fmt.Println("strconv failed with an error", err)
		return err
	}

	// do something else
	// ...

	return nil
}

func RunJob0(item string) error {
	if err := RunJob1(item); err != nil {
		fmt.Println("job1 failed with an error", err)
		return err
	}

	return nil
}

func main() {
	if err := RunJob0("   123   "); err != nil {
		fmt.Println("job0 failed with an error", err)
		return
	}
	fmt.Println("ok")
}
