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
		for {
			v, ok := <-unbuffered
			if !ok {
				fmt.Println("stop reader")
				return
			}
			fmt.Println(v)
		}
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
