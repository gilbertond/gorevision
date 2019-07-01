package main

import (
	"fmt"
	"strconv"
)

// No classes in Go but rather struct
// two method types (Value and pointer receivers)
//
type Person struct {
	firstName string
	lastName  string
	gender    string
	age       int
	city      string
}

func main() {
	person1 := Person{"jack", "", "", 20, ""}
	// Or
	person2 := Person{firstName: "peter", lastName: "", gender: "", age: 20, city: ""}

	fmt.Println(person1.firstName)
	fmt.Println(person2.firstName)

	fmt.Println(greetPerson(person1))

	//	update person
	person1.updateObject()
	fmt.Println(greetPerson(person1))
	person1.setLastName("Gilbert")
	fmt.Println(person1.firstName + " " + person1.lastName + " " + strconv.Itoa(person1.age))

}

//value receiver
func greetPerson(p Person) string {
	return "Hello " + p.firstName + " " + p.lastName + ", Your age is " + strconv.Itoa(p.age)
}

//pointer receiver (if you wanna change object value)
func (p *Person) updateObject() {
	p.firstName = "Gilbert"
}
func (p *Person) setLastName(fName string) {
	if p.firstName != fName {
		p.lastName = p.lastName
	} else {
		p.lastName = "Ndenzi"
	}
}

//Illustrate anonymous functions
func someFunction()  {

	//goroutine anonymous function
	func(msg string) {
		fmt.Println(msg)
	}("going")
}
