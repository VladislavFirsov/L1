package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	numbers := generateSlice()
	quickSort(numbers)
	fmt.Println(numbers)

}

func quickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	start, end := 0, len(arr)-1
	pivot := end
	for i := start; i < end; i++ {
		if arr[i] < arr[pivot] {
			arr[start], arr[i] = arr[i], arr[start]
			start++
		}

	}
	arr[start], arr[pivot] = arr[pivot], arr[start]
	quickSort(arr[:start])
	quickSort(arr[start+1:])
}

func generateSlice() []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, 10)

	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}
	return arr
}
