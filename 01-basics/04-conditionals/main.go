package main

import "fmt"

func main() {
	conditional_understanding()
}

func conditional_understanding() {
	const age = 18
	if age >= 18 {
		fmt.Println("You're an adult")
	} else {
		fmt.Println("You're not an adult")
	}
}
