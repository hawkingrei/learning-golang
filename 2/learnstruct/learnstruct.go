// learnstruct
package main

import (
	"fmt"
)

type Rect struct {
	x, y          float64
	width, height float64
}

func main() {
	rect1 := new(Rect)
	rect2 := &Rect{}
	rect3 := &Rect{0, 0, 100, 200}
	rect4 := &Rect{width: 100, height: 200}
	fmt.Println(rect1)
	fmt.Println(rect2)
	fmt.Println(rect3)
	fmt.Println(rect4)
}
