package main

import "fmt"

func main() {
	//Arrays
	var fruitArr [2]string
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	fruitArr2 := [2]string{"Banana", "Tomato"}

	fmt.Println(fruitArr[1])
	fmt.Println(fruitArr2[1])

	fruitSlice := []string{"Gurkin", "Mango"}

	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:2])
}
