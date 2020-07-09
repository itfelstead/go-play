package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	fileName := os.Args[1]
	h, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error", err.Error())
		os.Exit(1)
	}
	io.Copy(os.Stdout, h)
	h.Close()
}
