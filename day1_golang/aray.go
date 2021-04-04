package main

import "fmt"

func main() {
	arr := []int{2, 3, 4, 5, 6}
	for i, data := range arr {
		fmt.Println(i, " ", data)
	}
	///////
	arr1 := []int{2, 3, 4, 5, 6}
	for _, data := range arr1 {
		fmt.Println(" ", data)
	}
	////////////
	name := []string{"Rohit", "Pinky", "Sony", "Scooby"}
	for _, d := range name {
		fmt.Println(" ", d)
	}
}
