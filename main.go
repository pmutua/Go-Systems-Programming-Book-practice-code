package main

import "fmt"

func main() {

	// Slices
	aSlice := []int{-1, 4, 5, 7, 5, 4, 5, 6, 67}
	fmt.Printf("Before change: ")
	printSlice(aSlice)
	change(aSlice)
	fmt.Printf("After change: ")
	printSlice(aSlice)

	fmt.Printf("Before. Cap: %d, length: %d\n", cap(aSlice), len(aSlice))
	aSlice = append(aSlice, -100)
	fmt.Printf("After. Cap: %d, length: %d\n", cap(aSlice), len(aSlice))
	printSlice(aSlice)
	anotherSlice := make([]int, 4)
	fmt.Printf("A new slice with 4 elements: ")
	printSlice(anotherSlice)

	//Maps
	aMap := make(map[string]int)

	aMap["Mon"] = 0
	aMap["Tue"] = 1
	aMap["Wed"] = 2
	aMap["Thu"] = 3
	aMap["Fri"] = 4
	aMap["Sat"] = 5
	aMap["Sun"] = 6

	// Get Value of existing Key
	fmt.Printf("Sunday is the %dth day of the week.\n", aMap["Sun"])

}
