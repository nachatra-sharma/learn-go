package main

import "fmt"

func main() {
	arr := [5]int{6, 4, 17, 10, 12}
	fmt.Println(max(arr))
}

func max (arr [5]int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}
