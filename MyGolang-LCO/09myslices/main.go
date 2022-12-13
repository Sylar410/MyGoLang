package main

import (
	"fmt"
	"sort"
)

func main() {

	// Creating an array
	arr := [7]string{"welcome", "to", "slices", "This", "is", "an", "example"}
	fmt.Println("Array:", arr)
	myslice := arr[:7]
	fmt.Println("Slice:", myslice)
	fmt.Printf("Length of the slice: %d", len(myslice))
	fmt.Printf("\nCapacity of the slice: %d \n", cap(myslice))

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	//highScores[4] = 777

	highScores = append(highScores, 555, 666, 321)

	fmt.Println(highScores)

	//fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	//fmt.Println(highScores)

	//how to remove a value from slices based on index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
	for e := 0; e < len(myslice); e++ {
		fmt.Println(myslice[e])
	}
}
