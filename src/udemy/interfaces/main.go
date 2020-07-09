package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{} // no fields needed in this example
type spanishBot struct{} // no fields needed in this example

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb) // we can do this as eb satisfies 'bot's func requirements
	printGreeting(sb) // we can do this as sb satisfies 'bot's func requirements
}

func (englishBot) getGreeting() string {
	// <non generic code>
	return "Hello"
}

/*
func printGreeting(eb englishBot) {
	// <generic code>
	fmt.Println(eb.getGreeting())
}
*/

func (spanishBot) getGreeting() string {
	// <non generic code>
	return "Hola"
}

/*
func printGreeting(sb spanishBot) {
	// <generic code>
	fmt.Println(sb.getGreeting())
}
*/

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
