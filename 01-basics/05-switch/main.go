package main

import "fmt"

func main() {
	switch_understanding()
}

func switch_understanding() {
	var day = 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	default:
		fmt.Println("Good days")
	}
}
