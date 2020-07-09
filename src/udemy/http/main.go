package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type myLogWriter struct{} // no fields needed for now

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	fmt.Println(resp) // !! doesn't print structs' Body member because it is an Interface type

	// bs := make([]byte, 99999) // make a big slice
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// Alternative: use std lib  COpy call to copy from Reader to Writer
	//io.Copy(os.Stdout, resp.Body)

	// Using our own Writer
	lw := myLogWriter{}
	io.Copy(lw, resp.Body)

}

// Supporting this function makes myLogWriter a Writer type
func (myLogWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("byte count: ", len(bs))
	return len(bs), nil
}
