package main

import (
	"fmt"
	"sync"
)

type counter struct {
	mu    sync.Mutex
	digit int
}

func (c *counter) increment() {
	c.mu.Lock()
	c.digit++
	c.mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	c := &counter{}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			c.increment()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(c.digit)
}
