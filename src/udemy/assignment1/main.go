package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range a {
		var r string
		if v%2 == 0 {
			r = "even"
		} else {
			r = "odd"
		}
		fmt.Printf("%d is %s\n", v, r)
	}
}
