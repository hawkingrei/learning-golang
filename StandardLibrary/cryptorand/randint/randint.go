// randint
package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	var n *big.Int
	var err error

	max := *big.NewInt(99999999999)

	n, err = rand.Int(rand.Reader, &max)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%d\n", n)

}
