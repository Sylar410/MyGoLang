package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "Welcome to the input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Rating")

	//comma ok || comma error

	input, _ := reader.ReadString('\n')
	fmt.Println("thank you for rating", input)
	fmt.Printf("Type of this ratng %T", input)

}
