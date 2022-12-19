package main

import "fmt"

func main() {

	fmt.Println("Welcome to Structs")

	//NO INHERITANCE , SUPER , CHILD
	username := User{"Batman", 24, "Bm@mail", true}
	fmt.Println(username)
	username.GetStatus()
	username.NewEmail()
}

type User struct {
	Name   string
	Age    int
	email  string
	Status bool
}

func (u User) GetStatus() {
	fmt.Println("Status", u.Status)

}

func (u User) NewEmail() {
	u.email = "Ded@go"
	fmt.Println(u.email)
}
