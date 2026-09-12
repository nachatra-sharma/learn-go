package main

import "fmt"

func main() {
	fmt.Println("result: ", sum(1, 2, 3, 4))
}

// variadic function: accepts any number of int arguments
func sum(args ...int) int {
	sum := 0
	for _, value := range args {
		sum += value
	}
	return sum
}
