package main

import "fmt"

func main() {
	a := 1
	b := 10
	c := 14
	d := 48

	a, b = b, a
	changeValue(&c, &d)
	fmt.Println(a, b)
	fmt.Println(c, d)
}

func changeValue(a, b *int) {
	*a ^= *b // Using XOR
	*b ^= *a
	*a ^= *b
}
