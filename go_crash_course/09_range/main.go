package main

import "fmt"

func main(){
	ids := []int{33,72,23,56,123,2}

	for i, id := range ids {
		fmt.Printf( "%d - ID: %d\n", i, id)
	}

	for _, id := range ids {
		fmt.Printf("D: %d\n", id)
	}

	sum := 0
	for _, id := range ids {
		sum+=id
	}
	fmt.Println("Sum", sum)

	emails := map[string]string{"Bob": "Bob@gmail.com", "Karen": "Karen@gmail.com", "Tim": "Tim.gmail.com"}
	
	for k,v := range emails {
		fmt.Printf("%s: %s\n", k,v)
	}

	for k := range emails {
		fmt.Println("Name: " + k)
	}
}