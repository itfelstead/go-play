package main

import "fmt"

func main() {

	// var colours map[string]string
	// colours := make(map[string]string)

	colour := map[string]string{ /// map[<keytype>]<valuetype>
		"red":   "ff0000",
		"white": "ffffff",
		"green": "00ff00",
	}

	//colour["black"] = "000000"
	//delete(colour, "black")

	//fmt.Println(colour)

	printMap(colour)
}

func printMap(c map[string]string) {
	for colorKey, hexValue := range c {
		fmt.Println("key: ", colorKey, "is", hexValue)
	}
}
