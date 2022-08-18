package main

import (
	"fmt"
	"alpk.work/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Print(message)
}
