package main

import (
	"fmt"
	"sync"
)

func main() {
	unbuffered := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range unbuffered {
			fmt.Println(v)
		}
		fmt.Println("stop reader")
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i <= 9; i++ {
			unbuffered <- fmt.Sprintf("Hello #%d", i)
		}
		close(unbuffered)
	}()
	wg.Wait()
}
