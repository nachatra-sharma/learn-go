package main

import "fmt"

type User struct {
	Name string
	Age int
	Email string
}

func main() {
	myUser := User{
		Name: "Aman Sharma",
		Age: 17,
		Email: "amansharma@gmail.com",
	}

	fmt.Println("name: ", myUser.Name)
	fmt.Println("age: ", myUser.Age)
	fmt.Println("email: ", myUser.Email)
}
