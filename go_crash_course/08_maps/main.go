package main

import "fmt"

func main() {

	/* 	emails := make(map[string]string) */

	/* 	emails["Bob"] = "Bob@gmail.com"
	   	emails["Karen"] = "Karen@gmail.com"
	   	emails["Tim"] = "Tim@gmail.com" */

	emails := map[string]string{"Bob": "Bob@gmail.com", "Karen": "Karen@gmail.com", "Tim": "Tim.gmail.com"}

	fmt.Println(emails)
	fmt.Println(emails["Bob"])
	fmt.Println(len(emails))

	delete(emails, "Tim")
	fmt.Println(emails)
}
