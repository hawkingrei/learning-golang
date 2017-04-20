
package main

import (
	"fmt"
	"os"
)

func main() {
	var originalWD, _ = os.Getwd()
        fmt.Println(originalWD)
}
