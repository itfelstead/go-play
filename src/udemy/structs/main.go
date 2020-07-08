package main

import "fmt"

// embedded struct for jim
type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	//contact   contactInfo
	contactInfo
}

func main() {

	// alex := person{firstName: "explicitly", lastName: "labelled"}
	// alex := person{"assumes", "order"}
	//fmt.Println(alex)
	//fmt.Printf("%+v",alex")

	//embedded
	jim := person{
		firstName: "Jim",
		lastName:  "x",
		//contact: contactInfo{
		contactInfo: contactInfo{
			email:   "x@xx",
			zipCode: 12345,
		},
	}
	// fmt.Printf("%+v", jim)
	jim.printMe()

	jimPtr := &jim
	jimPtr.updateName("pass by ref")
	jim.printMe()
	// jimPtr.printMe() // also seems to work..
}

func (p *person) updateName(newName string) {
	//(*p).firstName = newName
	// p.firstName = newName // also seems to work..
}

func (p person) printMe() {
	fmt.Printf("%+v\n", p)
}
