package main

import "fmt"

func main() {
	fmt.Println("welcome to array")

	var numbers [4]int
	numbers[0] = 12
	numbers[1] = 43
	numbers[2] = 45

	fmt.Println(numbers)
	fmt.Println(len(numbers))

	var nums = [5]int{34, 67, 78}
	fmt.Println(nums)
	fmt.Println(len(nums))

	var num [5]int
	num[1] = 07
	num[3] = 9
	num[4] = 54
	fmt.Println(num)
	fmt.Println(len(num))

}
