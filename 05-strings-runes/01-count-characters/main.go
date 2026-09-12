package main

import "fmt"

func main() {
	s := "hello"
	countChar(s)
}

func countChar(s string) {
	mpp := make(map[rune] int)
	for _, char := range s {
		mpp[char]++
	}

	for char, value := range mpp {
		fmt.Printf("%c ---> %d\n", char, value)
	}
}
