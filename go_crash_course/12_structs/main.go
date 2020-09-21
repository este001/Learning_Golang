package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

//value reciver
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

//pointer recivr seter
func (p *Person) hasBirthday() {
	p.age++
}

func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}

}

func main() {
	person1 := Person{firstName: "Estefan", lastName: "Rengifo Marin", city: "Stockholm", gender: "f", age: 12}
	person2 := Person{"Manolo", "Garcia", "Stockholm", "male", 12}

	fmt.Println(person1)
	fmt.Println(person1.firstName)
	fmt.Println(person1.greet())
	fmt.Println(person1.age)
	person1.hasBirthday()
	fmt.Println(person1.age)
	person1.getMarried("Svensson")
	person2.getMarried("Andersson")
	fmt.Println(person1.greet())

}
