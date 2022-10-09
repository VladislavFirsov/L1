package main

import (
	"fmt"
	"sync"
	"time"
)

type dict struct {
	mu      sync.Mutex // mutex protects values
	hashMap map[string]int
}

func (d *dict) Increment(key string) {
	d.mu.Lock()
	d.hashMap[key]++
	d.mu.Unlock()
}

func (d *dict) GetValue(key string) int {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.hashMap[key]
}

func main() {
	key := "number"
	d := dict{hashMap: make(map[string]int)}
	for i := 0; i < 100; i++ {
		go d.Increment(key)
	}

	time.Sleep(1 * time.Second)
	fmt.Println(d.GetValue(key))
}
