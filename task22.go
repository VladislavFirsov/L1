package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var a, b = rand.Int63(), rand.Int63()
	answer1 := calculator(a, b, "+")
	answer2 := calculator(a, b, "-")
	answer3 := calculator(a, b, "*")
	answer4 := calculator(a, b, "/")
	fmt.Printf("a + b = %v\n a - b = %v\n a * b = %v\n a / b = %v\n", answer1, answer2, answer3, answer4)
}

func calculator(a, b int64, c string) *big.Int {
	result := big.NewInt(0)
	switch c {
	case "+":
		result.Add(big.NewInt(a), big.NewInt(b))
	case "-":
		result.Div(big.NewInt(a), big.NewInt(b))
	case "*":
		result.Mul(big.NewInt(a), big.NewInt(b))
	case "/":
		result.Div(big.NewInt(a), big.NewInt(b)) // cannot put 0 in the second value, will panic
	}
	return result
}
