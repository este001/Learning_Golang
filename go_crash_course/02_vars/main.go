package main

import "fmt"

var name string = "este"

//name := "este"    går ej utanför

func main() {

	var age = 37
	const isCool = true
	size := 1.3

	tele, email := 070717034, "este@mail.com"

	fmt.Println(name, age, isCool, size, tele, email)
	fmt.Printf("%T\n", age)
}
