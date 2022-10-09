package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go stopByCancel(&wg)
	go stopByChannel(&wg)
	wg.Wait()

}

func stopByCancel(wg *sync.WaitGroup) {
	ctx, cancel := context.WithCancel(context.Background()) //similar works with context.WithDeadline and context.WithTimeout
	defer wg.Done()
	defer cancel()
	go func() {
		fmt.Println("from stopByCancel")
		cancel()

	}()
	time.Sleep(500 * time.Millisecond)
	<-ctx.Done()
}

func stopByChannel(wg *sync.WaitGroup) {
	defer wg.Done()
	done := make(chan bool)
	go func() {
		fmt.Println("from stopByChannel")
		done <- true

	}()
	<-done

}
