package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions")
	greeting()
	result := add(6, 9)
	fmt.Println(result)
	proRes, message := proAdd(5, 7, 5, 4, 4, 4, 7)
	fmt.Println(proRes)
	fmt.Println(message)

}

func add(valOne int, valTwo int) int {
	return valOne + valTwo
}

func greeting() {
	fmt.Println("Hello")
}

func proAdd(values ...int) (int, string) {
	total := 0
	for _, v := range values {
		total += v
	}
	return total, "Hello World"
}
