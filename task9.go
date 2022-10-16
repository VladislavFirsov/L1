package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	numbers := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numChan := make(chan int, 10)
	squareChan := make(chan int, 10)
	doneNumChan := make(chan bool)
	doneSquareChan := make(chan bool)
	wg.Add(1)
	go transferFromListToNumChan(&wg, numbers, numChan)
	go transferFromNumChanToSquareChan(numChan, squareChan, doneNumChan)
	go printFromSquareChan(squareChan, doneSquareChan)
	wg.Wait()
	close(numChan)
	<-doneNumChan
	close(squareChan)
	<-doneSquareChan
}

func transferFromListToNumChan(wg *sync.WaitGroup, numbers [10]int, numChan chan<- int) {
	defer wg.Done()
	for _, value := range numbers {
		numChan <- value
	}

}

func transferFromNumChanToSquareChan(numChan <-chan int, squareChan chan<- int, doneNumChan chan<- bool) {
	for value := range numChan {
		squareChan <- value * value
	}
	doneNumChan <- true
}

func printFromSquareChan(squareChan <-chan int, doneSquareChan chan<- bool) {
	for value := range squareChan {
		fmt.Println(value)
	}
	doneSquareChan <- true
}
