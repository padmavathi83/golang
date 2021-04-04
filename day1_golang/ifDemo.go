package main

import "fmt"

func main() {
	age := 20
	if age >= 20 {
		fmt.Println("Person can vote")
	} else if age < 20 {
		fmt.Println("Person cant vote")
	} else {
		fmt.Println("Invalid age")
	}
}


//accept 3 variable values from user and check which value is bigger