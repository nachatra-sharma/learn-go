package main

import "fmt"

func main() {
	a := 2
	b := 3
	swap(&a, &b)
	fmt.Println("Value of a is: ", a)
	fmt.Println("Value of b is: ", b)
}

func swap(a *int, b *int) {
	*a, *b = *b, *a
}

// func change(x *int) {
//     *x = 100
// }

// func main() {
//     x := 10
//     change(&x)

//     fmt.Println(x) // 100
// }

// func change(x *int) {
//     *x = 12
// }

// func main() {
//     x := 10
//     p := &x

//     fmt.Println(p)

//     change(p)

//     fmt.Println(*p)
// }
