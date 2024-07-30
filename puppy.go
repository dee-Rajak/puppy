package puppy

import (
	"fmt"

	"github.com/dee-Rajak/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof Woof Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func From10() {
	fmt.Println("I am from version 1.0.0")
}

func From11() {
	fmt.Println("I am from version 1.1.0")
}

func From12() {
	fmt.Println("I am from version 1.2.0")
}
