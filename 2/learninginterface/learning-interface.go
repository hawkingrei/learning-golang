// learning-interface
package main

import (
	"fmt"
)

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}
func (a *Integer) Add(b Integer) {
	*a += b
}

type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}

func main() {
	var a Integer = 1
	var b LessAdder = &a
	b.Add(2)
	fmt.Println(a)
}
