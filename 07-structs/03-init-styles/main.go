package main

import "fmt"

type User struct {
	Name string
	Age int
	Email string
	isAdult bool
}

func main() {
	var user User
	myUser := User{"Rishi", 23, "rishi@gmail.com", true}
	newUser := User{
		Name: "Rahul",
		Age: 20,
		Email: "rahul@gmail.com",
		isAdult: true,
	}
	fmt.Println(user)
	fmt.Println(myUser)
	fmt.Println(newUser)
}
