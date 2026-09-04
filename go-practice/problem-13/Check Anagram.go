package main

import "fmt"

func main() {
	s := "listen"
	t := "silent"
	fmt.Println(isAnagram(s, t))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	smpp := make(map[rune] int)

	for _, value := range s {
		smpp[value]++
	}

	for _, value := range t {
		if smpp[value] != 0 {
			smpp[value]--
		} else if smpp[value] == 0 {
			return false
		}
	}
	return true
}
