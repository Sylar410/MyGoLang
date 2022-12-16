package main

import "fmt"

func main() {

	fmt.Println("Welcome to Structs")

	//NO INHERITANCE , SUPER , CHILD
	username := User{"Batman", 24, "Bm@mail"}
	fmt.Println(username)

}

type User struct {
	Name  string
	Age   int
	email string
}
