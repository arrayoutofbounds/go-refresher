package main

import (
	"fmt"
	"strconv"
)

// Person is a person
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// value receiver. Just for doing calculations and not changing anything
func (p Person) greet() string {
	return "Hello my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// pointer receivers are for changing things
// methods that cause change
func (p *Person) hasBirthday() {
	p.age++
}

// get married pointer receiver
func (p *Person) getMarried(spouselastname string) {
	if p.gender == "M" {
		return
	} else {
		p.lastName = spouselastname
	}
}

func main() {
	person1 := Person{
		firstName: "Sam",
		lastName:  "Smith",
		city:      "sydney",
		gender:    "F",
		age:       25}

	// or
	// person1 := {"Sam", "Smith", "Sydney", "M", 25}

	//fmt.Println(person1)
	//fmt.Println(person1.firstName)

	person1.hasBirthday()
	person1.getMarried("williams")
	fmt.Println(person1.greet())

}
