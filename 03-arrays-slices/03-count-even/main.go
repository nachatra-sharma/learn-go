package main

import "fmt"


func main() {
	arr := [5]int{1, 2, 3, 4, 6}
	fmt.Println(countEven(arr))
}


func countEven(arr [5]int) int {
	counter := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] % 2 == 0{
			counter++
		}
	}
	return counter
}
