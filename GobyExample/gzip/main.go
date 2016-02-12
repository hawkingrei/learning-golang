package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
)

func main() {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write([]byte("YourDataHere")); err != nil {
		panic(err)
	}
	if err := gz.Flush(); err != nil {
		panic(err)
	}
	if err := gz.Close(); err != nil {
		panic(err)
	}
	fmt.Println(b)
}
