package main

import "fmt"


func main () {
	arr := []int{1, 2, 3, 4, 5, 6}
	target := 15
	fmt.Println(findElement(arr, target))
}


func findElement(arr []int, target int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return true
		}
	}
	return false
}
