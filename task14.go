package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a, b, c, d interface{} = "hello", 1, true, make(chan int)
	defineTypeByFormat(c)
	defineTypeBySwitch(a)
	defineTypeBySwitch(b)
	defineTypeByReflect(d)

}
func defineTypeByFormat(v interface{}) {
	fmt.Printf("Type of your value is: %T\n", v)

}

func defineTypeByReflect(v interface{}) {
	fmt.Println("You've got", reflect.TypeOf(v))
}

func defineTypeBySwitch(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("This is int")
	case string:
		fmt.Println("This is string")
	case bool:
		fmt.Println("This is bool")
	default:
		fmt.Println("This function doesn't know this type, try something else")
	}
}
