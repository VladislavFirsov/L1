package main

import "fmt"

func main() {
	temperatures := [8]float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	tempGroups := make(map[int][]float32)
	for _, value := range temperatures {
		key := int(value) / 10 * 10
		if _, ok := tempGroups[key]; !ok {
			tempGroups[key] = make([]float32, 0)
		}
		tempGroups[key] = append(tempGroups[key], value)
	}
	fmt.Println(tempGroups)
}
