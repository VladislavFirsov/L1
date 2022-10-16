package main

func main() {
	//Originally, after a string is created it has a pointer to an immutable massive of BYTES
	//but if we are working with letters, we are likely to use type rune instead in order not to
	//get runtime panic, so to make it work we have to convert bytes to runes.

	//var justString string
	//func someFunc(){
	// v :=createHugeString(1 << 10)
	// allRunes := []runes(v)
	// justString = allRunes[:100]

	//}
}
