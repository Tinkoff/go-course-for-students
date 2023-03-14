package main

import "fmt"

func main() {
	unbuffered := make(chan string)
	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			v, ok := <-unbuffered
			if !ok {
				fmt.Println("stop reader")
				return
			}
			fmt.Println(v)
		}
	}()

	go func() {
		for i := 0; i <= 9; i++ {
			unbuffered <- fmt.Sprintf("Hello #%d", i)
		}
		close(unbuffered)
	}()

	<-done

	fmt.Println("done")
}
