package main

import "fmt"

func greeting(name string) string {
	return "hello" + name
}

func getSum(num1, num2 int) int {

	return num1 + num2

}

func main() {
	const name = "Tim"
	fmt.Println(greeting(name))
	fmt.Println(getSum(1, 2))
}
