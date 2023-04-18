package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	eg, ctx := errgroup.WithContext(context.Background()) // ~ context.WithCancel() + waitGroup

	eg.Go(func() error {
		return RunWorkers(ctx)
	})

	//go RunWorkers(ctx)

	sigQuit := make(chan os.Signal, 1)
	signal.Notify(sigQuit, syscall.SIGINT, syscall.SIGTERM)

	eg.Go(func() error {
		select {
		case s := <-sigQuit:
			return fmt.Errorf("captured signal: %v", s)
		case <-ctx.Done():
			return nil
		}
	})

	if err := eg.Wait(); err != nil {
		fmt.Printf("gracefully shutting down the server: %s\n", err.Error())
	}

	fmt.Println("graceful shutdown service")
}

// bad code. only for mock
func generator() chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		i := 0
		for {
			out <- i
			time.Sleep(time.Second)
			i += 1
		}
	}()

	return out
}

func RunWorkers(ctx context.Context) error {
	wn1Chan := workerNode1(ctx)
	return workerNode2(wn1Chan)
}

func workerNode1(ctx context.Context) chan string {
	input := generator()
	out := make(chan string)

	go func() {
		defer func() {
			fmt.Println("workerNode1 stop")
			close(out)
		}()

		for {
			select {
			case <-ctx.Done():
				return
			case msg := <-input:
				out <- strconv.Itoa(msg)
			}
		}
	}()

	return out
}

func workerNode2(input chan string) error {
	for msg := range input {
		fmt.Println(msg)
	}

	fmt.Println("workerNode2 stop")
	return nil
}
