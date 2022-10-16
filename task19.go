package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	reversedStr := reverseString(s)
	fmt.Println(reversedStr)

}

func reverseString(s string) string {
	r := []rune(s)
	for start, end := 0, len(r)-1; start < end; start, end = start+1, end-1 {
		r[start], r[end] = r[end], r[start]
	}
	return string(r)
}
