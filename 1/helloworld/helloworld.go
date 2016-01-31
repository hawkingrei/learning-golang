// helloworld
package main

import "fmt"

func main() {
	var a int = 1
	m := 10
	m = m + 1
	const (
		Cyan    = iota
		Megenta = iota
	)
	var b *int = &a
	var c **int = &b
	var x int = *b
	fmt.Println(Cyan)
	fmt.Println(Megenta)
	fmt.Println("hello" + "world")
	fmt.Println("hello", "world")
	fmt.Println("a = ", a)
	fmt.Println("&a = ", &a)
	fmt.Println("*&a = ", *&a)
	fmt.Println("b = ", b)
	fmt.Println("&b = ", &b)
	fmt.Println("*&b = ", *&b)
	fmt.Println("*b = ", *b)
	fmt.Println("c = ", c)
	fmt.Println("*c = ", *c)
	fmt.Println("&c = ", &c)
	fmt.Println("*&c = ", *&c)
	fmt.Println("**c = ", **c)
	fmt.Println("***&*&*&*&c = ", ***&*&*&*&*&c)
	fmt.Println("x = ", x)
}
