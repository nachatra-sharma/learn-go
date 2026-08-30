package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 3, 2, 1}
	findFrequency(arr)
	for key, value := range arr {
		fmt.Println(key, value)
	}
}


func findFrequency (arr []int) {
	frequency := make(map[int]int)
	for _, value := range arr {
		frequency[value]++
	}
	fmt.Println(frequency)
}
