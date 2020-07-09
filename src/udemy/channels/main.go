package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	//for i := 0; i < len(links); i++ {
	//for { // infinite loop
	//	go checkLink(<-c, c)
	//fmt.Println(<-c) //block waiting for message
	//}

	for l := range c {
		//time.Sleep(time.Second * 5)
		//go checkLink(l, c)
		go func(mylink string) {
			time.Sleep(time.Second * 5)
			checkLink(mylink, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	}
	fmt.Println(link, "is up")
	c <- link
}
