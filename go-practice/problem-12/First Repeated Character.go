package main

import "fmt"

func main() {
	s := "abcdbea"
	fmt.Printf("The repeating char is %c\n",findFirstRepeatingChar(s))
}

func findFirstRepeatingChar(s string) rune {
	mpp := make(map[rune] int)
	var answer rune

	for _, value := range s {
		mpp[value]++
	}

	for _, value := range s {
		if mpp[value] == 2 {
			answer = value
		}
	}
	return answer
}
