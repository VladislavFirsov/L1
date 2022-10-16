package main

import (
	"fmt"
)

func main() {
	var (
		num, pos int64
		bit      uint8
	)
	num = 7 // any int64
	pos = 1 // starts from the right
	bit = 0 // 1 or 0
	fmt.Println(setBit(num, pos, bit))

}

func setBit(num int64, pos int64, bit uint8) int64 {
	switch bit {
	case 1:
		num = num | (1 << pos)

	case 0:
		num = num & ^(1 << pos)

	}
	return num
}
