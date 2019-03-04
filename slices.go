package main

import "fmt"

func change(x []int) {
	x[3] = -2
}

func printSlice(x []int) {
	for _, number := range x {
		fmt.Printf("%d ", number)
	}
	fmt.Println()
}
