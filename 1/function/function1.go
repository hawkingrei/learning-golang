// function1
package main

import (
	"errors"
	"fmt"
)

func Add(a int, b int) (ret int, err error) {
	if a < 0 || b < 0 {
		err = errors.New("Should be non-negative numbers!")
		return
	}
	return a + b, nil
}
func main() {
	value, _ := Add(10, 12)
	fmt.Println(value)
}
