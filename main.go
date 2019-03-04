package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type message struct {
	X     int
	Y     int
	Label string
}

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

	// Error Handling
	_, ok := aMap["Tuesday"]
	if ok {
		fmt.Printf("the Tuesday key exists!\n")
	} else {
		fmt.Printf("The Tuesday key does not exist!\n")
	}

	// Iterating over all the Keys of an existing map

	count := 0
	for key, _ := range aMap {
		count++
		fmt.Printf("%s ", key)
	}
	fmt.Printf("\n")
	fmt.Printf("The aMap has %d elements\n", count)

	// Alternative way of initializing a map

	anotherMap := map[string]int{
		"One":  1,
		"Two":  2,
		"Tree": 3,
		"Four": 4,
	}

	anotherMap["Five"] = 5
	count = 0
	for _, _ = range anotherMap {
		count++
	}
	fmt.Printf("anotherMap had %d elements\n", count)

	// Convert Array to map
	anArray := [4]int{1, 3, 4, 5}
	aMap2 := make(map[string]int)

	length := len(anArray)

	for i := 0; i < length; i++ {
		fmt.Printf("%s ", strconv.Itoa(i))
		aMap2[strconv.Itoa(i)] = anArray[i]
	}

	for key, value := range aMap {
		fmt.Printf("%s: %d\n", key, value)
	}

	// Go data structures

	p1 := message{23, 12, "A message"}
	p2 := message{}
	p2.Label = "Message 2"

	s1 := reflect.ValueOf(&p1).Elem()
	s2 := reflect.ValueOf(&p2).Elem()
	fmt.Println("S2= ", s2)

	// Print all fields without knowing thier names
	typeOfT := s1.Type()
	fmt.Println("P1=", p1)
	fmt.Println("P2=", p2)

	for i := 0; i < s1.NumField(); i++ {
		f := s1.Field(i)

		fmt.Printf("%d: %s ", i, typeOfT.Field(i).Name)
		fmt.Printf("%s = %v\n", f.Type(), f.Interface())
	}

}
