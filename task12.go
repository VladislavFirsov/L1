package main

import "fmt"

func main() {
	someStrings := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(Set(someStrings))
}

func Set(s []string) []string {
	setMap := make(map[string]int)
	set := make([]string, 0)
	for _, value := range s {
		if setMap[value] >= 1 {
			continue
		}

		setMap[value]++

	}
	for key, value := range setMap {
		if value == 1 {
			set = append(set, key)
		}
	}
	return set
}
