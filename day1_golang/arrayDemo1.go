package main

import "fmt"

func main() {
	grade1 := 98
	grade2 := 87
	grade3 := 85
	fmt.Printf("Grades: %v, %v, %v\n", grade1, grade2, grade3)
	//
	grades := [3]int{98, 87, 85}
	fmt.Printf("grades:%v", grades)
}
