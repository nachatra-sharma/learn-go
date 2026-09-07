package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	for i, char := range "Nachatra" {
		fmt.Println(i, char)
	}
	variables()
	simpleValues()
	loop_understanding()
	conditional_understanding()
	switch_understanding()
	fmt.Println("result: ", sum(1,2,3,4))
}

func simpleValues() {
	fmt.Println("Nachatra Sharma")
	fmt.Println(123)
	fmt.Println(3.14)
	fmt.Println(true)
	fmt.Printf("%c\n", 65)
}

func variables() {
	var productName string = "Iphone 16 Pro Max"
	age := 23
	const PI_VALUE = 3.14
	fmt.Println(age)
	fmt.Println(productName)
	fmt.Println(PI_VALUE)
}

func loop_understanding() {
	// for loop
	for i := 1; i <= 5; i++ {
		fmt.Println("The value of i is: ", i);
	}

	// while loop using for loop
	j := 1;
	for j <= 5 {
		fmt.Println("The value of 2 * j is: ", 2 * j)
		j++;
	}
}

func conditional_understanding() {
	const age = 18
	if age >= 18 {
		fmt.Println("You're an adult")
	} else {
		fmt.Println("You're not an adult")
	}
}

func switch_understanding() {
	var day = 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	default:
		fmt.Println("Good days")
	}
}

func sum(args ...int) int{
	sum := 0
	for _, value := range args {
		sum += value
	}
	return sum
}
