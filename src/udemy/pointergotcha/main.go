package main

import "fmt"

func main() {
	myslice := []string{"a", "b", "c"}

	attemptChange(myslice)

	fmt.Println(myslice)

}

func attemptChange(s []string) {
	s[0] = "must pass by ref"
}
