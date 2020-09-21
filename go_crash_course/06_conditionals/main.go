package main

import "fmt"

func main(){
	x := 10
	y := 5

	if x < y {
		fmt.Printf("%d is less than %d\n", x,y)
	} else {
		fmt.Printf("%d is less than %d\n", y,x)
	}

	color := "red"

	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println("other color")
	}

	switch color {
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("color is blue")
	default:
		fmt.Println("other color")
	}
}