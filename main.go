package main

import "fmt"

func main() {
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

}
