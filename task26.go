package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aabcd"
	fmt.Println(isUnique(str1))
	fmt.Println(isUnique(str2))
	fmt.Println(isUnique(str3))
}

func isUnique(s string) bool {
	runes := []rune(strings.ToLower(s))
	counter := make(map[rune]int)
	for _, value := range runes {
		_, ok := counter[value]
		if ok {
			return false
		}
		counter[value]++
	}
	return true
}
