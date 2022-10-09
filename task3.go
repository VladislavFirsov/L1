package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := [5]int{2, 4, 6, 8, 10}
	squares := make(chan int)
	done := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(len(numbers))
	for i := range numbers {
		go makesquare(&wg, squares, numbers[i])
	}

	go makesum(squares, done)
	wg.Wait()

	close(squares)
	<-done
}
func makesquare(wg *sync.WaitGroup, sq chan int, el int) {
	defer wg.Done()
	sq <- el * el
}

func makesum(sq chan int, done chan bool) {
	sum := 0
	for i := range sq {
		sum += i
	}

	fmt.Println(sum)
	done <- true
}
