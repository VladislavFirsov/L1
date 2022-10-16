package main

import "fmt"

func main() {
	numbers := []int{15, 13, 28, 7, 666, 534, 50}
	QuickSort(numbers)
	fmt.Println(binarySearch(numbers, 666))

}

func binarySearch(arr []int, search int) bool {
	start, end := 0, len(arr)-1

	for start <= end {
		mid := (start + end) / 2
		switch {
		case arr[mid] == search:
			return true
		case arr[mid] < search:
			start = mid + 1
		case arr[mid] > search:
			end = mid - 1
		}

	}
	return false
}

func QuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	start, end := 0, len(arr)-1
	pivot := end
	for i := 0; i < end; i++ {
		if arr[i] < arr[pivot] {
			arr[start], arr[i] = arr[i], arr[start]
			start++
		}
	}
	arr[start], arr[pivot] = arr[pivot], arr[start]
	QuickSort(arr[:start])
	QuickSort(arr[start+1:])
}
