package main

import "fmt"

func main() {
	x := 2
	p := &x
	fmt.Println(x, p, *p)
	fmt.Println("are they same ? ", &x, p)
	fun()
	fmt.Println("Before Function Call: ", *p)
	changeValueOfX(p)
	fmt.Println("After function call: ", *p)
	dryRun()
	changeXandY()
	changeTheValueOfX()
	checkIfBothAreEqual()
}

func changeValueOfX(p *int) {
	*p = 50
	fmt.Println("checking iiiiiii",*p)
}

/**
Explain this to me
x ---> it is a normal variable which hold some type of value
p ---> it hold the address of x
*p ---> it give us back the value of the address here p must have be an address
&x ---> it give us the address of x where value x value is stored
**/
func fun() {
	x := 25
	p := &x
	// print value of x using only p
	fmt.Println(*p)
}


func dryRun() {
	x := 10
	p := &x

	fmt.Println(x) // 10
	fmt.Println(*p) // 10

	*p = 20

	fmt.Println(x) // 20
	fmt.Println(*p) // 20
}


func changeXandY() {
	x := 10
	y := 20
	p1 := &x
	p2 := &y
	*p1 = 100
	*p2 = 200
	fmt.Println(*p1, *p2)
}

func changeTheValueOfX(){
	x := 10
	p := &x
	q := &x
	*q = 50
	if p == q {
		fmt.Println("Both are equal")
	} else {
		fmt.Println("Both are unequal")
	}
	fmt.Println(*q, p)
}

func checkIfBothAreEqual() {
	x := 10
	y := 10

	p := &x
	q := &y
	fmt.Println(p == q)
	fmt.Println(*p == *q)
}
