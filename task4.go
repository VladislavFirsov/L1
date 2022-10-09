package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	workers := runtime.NumCPU() // optimal usage of workers is to get maximum of Computer's cores we've got
	input := make(chan any, workers)
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGINT)

	go func() {
		<-sigChan
		fmt.Println("cancelling the program")
		close(input)
		close(sigChan)
		cancel()
		os.Exit(0)
	}()

	for i := 0; i < workers; i++ {
		go readFromInput(ctx, input)
	}

	for {
		num := rand.Intn(100)
		input <- num
		time.Sleep(100 * time.Millisecond)
	}

}

func readFromInput(ctx context.Context, input chan any) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("cancelling process")
			return
		case value := <-input:
			fmt.Printf("recieved value from input: %v\n", value)
		}
	}
}
