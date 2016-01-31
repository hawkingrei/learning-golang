// slice2
package main

import (
	"fmt"
)

func main() {
	mySlice := make([]int, 5, 10)
	fmt.Println("len(mySlice): ", len(mySlice))
	fmt.Println("cap(mySlice): ", cap(mySlice))
	mySlice2 := append(mySlice, 8, 9, 10)
	fmt.Println("=================================")
	for i := 0; i < len(mySlice2); i++ {
		fmt.Println("mySlice2[", i, "] =", mySlice2[i])
	}
	mySlice3 := append(mySlice, mySlice2...)
	fmt.Println("=================================")
	for i := 0; i < len(mySlice3); i++ {
		fmt.Println("mySlice2[", i, "] =", mySlice3[i])
	}
	fmt.Println("=================================")
	oldSlice := []int{1, 2, 3, 4, 5}
	newSlice := oldSlice[:3]
	for _, v := range newSlice {
		fmt.Print(v, " ")
	}
	fmt.Println("\n=================================")
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	copy(slice2, slice1)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println("---------------------------------")
	copy(slice1, slice2)
	fmt.Println(slice1)
	fmt.Println(slice2)
}
