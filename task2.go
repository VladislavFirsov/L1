package main

import "fmt"

func main() {
	numbers := [5]int{2, 4, 6, 8, 10}
	squares := make(chan int, 5)
	done := make(chan bool)
	go func(numbers [5]int, sq chan int) {
		for _, i := range numbers {
			sq <- i * i
		}
	}(numbers, squares)
	go func(sq chan int, d chan bool) {
		for i := 0; i < 5; i++ {
			fmt.Println(<-sq)
		}
		done <- true
	}(squares, done)

	<-done
}
