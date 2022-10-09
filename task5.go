package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	var wg sync.WaitGroup
	input := make(chan any)
	wg.Add(2)

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Finished by time")
				wg.Done()
				return
			case <-input:
				fmt.Println(<-input)
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()
	go func() {

		defer close(input)
		for {
			value := rand.Intn(100)
			select {
			case <-ctx.Done():
				wg.Done()
				return
			case input <- value:
			}

		}
	}()
	wg.Wait()

}
