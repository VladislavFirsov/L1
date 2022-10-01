package main

import "fmt"

type Human struct {
	name string
}

func (h *Human) ChangeName(name string) {
	h.name = name
}

type Action struct {
	name string
	Human
}

func main() {
	person := Human{
		name: "Andrew", //person name is Andrew
	}
	act := Action{
		name: "Action", // Struct name is action
		Human: Human{
			"Vasya", // Human name in the structure "Action" is Vasya
		},
	}

	person.ChangeName("Igor") // Changes the person's name
	act.ChangeName("Anton")   // Changes the name of Action structure

	fmt.Println(person.name)    // Igor
	fmt.Println(act.name)       // NewAction
	fmt.Println(act.Human.name) // Anton

}
