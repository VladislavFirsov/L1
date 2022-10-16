package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6}
	slice2 := []int{10, 11, 3}
	showIntersection(slice1, slice2)

}

func showIntersection(slice1, slice2 []int) {
	intersectionMap := make(map[int]int)
	intersection := make([]int, 0)
	for _, value := range slice1 {
		intersectionMap[value]++
	}

	for _, value := range slice2 {
		_, ok := intersectionMap[value]
		if ok {
			intersection = append(intersection, value)
		}
	}
	fmt.Println(intersection)
}
