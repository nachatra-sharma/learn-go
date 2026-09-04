package main

import "fmt"

type User struct {
	Name string
	Age int
}

func brandNewUser(name string, age int) *User {
	createdUser := User{
		Name: name,
		Age: age,
	}
	return &createdUser
}

func main() {
	user1 := brandNewUser("Rishi", 23)
	fmt.Println(user1.Name, user1.Age)
	user1.Name = "Nachatra Sharma"
	fmt.Println(user1.Name, user1.Age)
}
