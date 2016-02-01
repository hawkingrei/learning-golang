// less
package main

import (
	"fmt"
)

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}

//func (a *Integer) Add (b Integer){
//	*a += b
//}
func (a Integer) Add(b Integer) {
	a += b
}
func main() {
	var a Integer = 1
	if a.Less(2) {
		fmt.Println(a, "less 2")
	}
	a.Add(2)
	fmt.Println("a = ", a)
	fmt.Println("===============================================")
	var b = [3]int{1, 2, 3}
	var c = &b
	b[1]++
	fmt.Println(b, *c)
	fmt.Println("===============================================")
	d := b
	d[1]++
	fmt.Println(b, d)
	fmt.Println("===============================================")
}
