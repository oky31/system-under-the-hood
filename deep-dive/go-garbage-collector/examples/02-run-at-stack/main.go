package main

import "fmt"

type User struct {
	Id    int
	Name  string
	Email string
}

func CreateNewUser() User {
	return User{
		Id:    1,
		Name:  "Susilo",
		Email: "susilo@mail.com",
	}
}

func main() {
	user := CreateNewUser()
	fmt.Println(user)
}
