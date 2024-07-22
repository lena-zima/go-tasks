package main

import "fmt"

var Global = 5

func ChangeValue() {
	defer func(startValue int) {
		Global = startValue
	}(Global)
	Global = 3
	fmt.Println("New value: ", Global)
}

func main() {
	fmt.Println("Initial value: ", Global)
	ChangeValue()
	fmt.Println("Current value: ", Global)
}
