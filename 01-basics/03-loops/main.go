package main

import "fmt"

func main() {
	loop_understanding()
}

func loop_understanding() {
	// for loop
	for i := 1; i <= 5; i++ {
		fmt.Println("The value of i is: ", i)
	}

	// while loop using for loop
	j := 1
	for j <= 5 {
		fmt.Println("The value of 2 * j is: ", 2*j)
		j++
	}
}
