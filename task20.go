package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	str = strings.Trim(str, "\n\r")
	reversedSentence := reverseWordsInSentence(str)
	fmt.Println(reversedSentence)

}

func reverseWordsInSentence(s string) string {
	splittedStr := strings.Split(s, " ")
	for i := 0; i < len(splittedStr)/2; i++ {
		splittedStr[i], splittedStr[(len(splittedStr)-1)-i] = splittedStr[(len(splittedStr)-1)-i], splittedStr[i]
	}
	return strings.Join(splittedStr, " ")
}
