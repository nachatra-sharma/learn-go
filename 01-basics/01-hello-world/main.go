package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	// range over a string gives index + rune (character code)
	for i, char := range "Nachatra" {
		fmt.Println(i, char)
	}

	simpleValues()
}

func simpleValues() {
	fmt.Println("Nachatra Sharma")
	fmt.Println(123)
	fmt.Println(3.14)
	fmt.Println(true)
	fmt.Printf("%c\n", 65)
}
