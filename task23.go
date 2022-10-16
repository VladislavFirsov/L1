package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	slice := createSlice()
	fmt.Println(slice)
	deleteFromSlice(slice, 1)
	fmt.Println(slice)
}

func deleteFromSlice(slice []int, elem int) []int {
	return append(slice[:elem], slice[elem+1:]...)
}

func createSlice() []int {
	rand.Seed(time.Now().UnixNano())
	slice := make([]int, 10)
	for i := 0; i < 9; i++ {
		slice[i] = rand.Intn(100)
	}
	return slice
}
