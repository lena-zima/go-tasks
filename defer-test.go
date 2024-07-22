package defertest

import "fmt"

var Global = 5

func ChangeValue() {
	defer func(startValue int) {
		Global = startValue
	}(Global)
	Global = 3
	fmt.Println("New value: ", Global)
}

func defertest() {
	fmt.Println("Initial value: ", Global)
	ChangeValue()
	fmt.Println("Current value: ", Global)
}
