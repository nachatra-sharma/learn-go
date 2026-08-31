package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 3, 4}
	brr := []int{5, 4, 3, 3, 2, 2, 1}

	intersection(arr, brr)
}

func intersection(arr []int, brr []int) {
	marr := make(map[int] int)
	mbrr := make(map[int] int)
	intersectionArr := []int{}
	for _, value := range arr {
		marr[value]++;
	}
	for _, value := range brr {
		mbrr[value]++;
	}
	for key := range marr {
		if mbrr[key] != 0 {
			intersectionArr = append(intersectionArr, key)
		}
	}
	for _, value := range intersectionArr {
		fmt.Println(value)
	}
}
