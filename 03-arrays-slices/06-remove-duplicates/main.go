package main

import "fmt"

func main() {
	arr := []int{1, 2, 1, 4, 3, 3, 2}
	removeDuplicate(arr)
}

func removeDuplicate(arr []int) {
	result := []int{}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if (arr[i] == arr[j]) && (arr[i] != -1 || arr[j] != -1) {
				arr[j] = -1
			}
		}
	}

	for _, value := range arr {
		if value != -1 {
			result = append(result, value)
		}
	}

	for _, value := range result {
		fmt.Println(value)
	}
}
