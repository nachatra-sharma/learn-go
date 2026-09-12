package main

import "fmt"


func main() {
	arr := []int{1, 3, 4, 2, 8, 12, 3}
	fmt.Println(findDuplicate(arr))
}

func findDuplicate(arr []int) int {
	duplicate := -1
	frequency := make(map[int] int)
	for _, value := range arr {
		frequency[value]++
	}
	for key, value := range frequency {
		if value == 2 {
			duplicate = key
		}
	}
	return duplicate
}
