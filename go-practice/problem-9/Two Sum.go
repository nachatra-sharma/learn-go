package main

import "fmt"

func main() {
	arr := []int {2,7,11,15}
	target := 9
	fmt.Println(findIndex(arr, target))
}


func findIndex(arr []int, target int) []int {
	frequency := make(map[int]int)

	for key, value := range arr {

		needed := target - value

		if frequency[needed] != 0 {
			return []int{frequency[needed] - 1, key}
		}

		frequency[value] = key + 1
	}

	return []int{}
}
