package main

import "fmt"

func main(){
	fmt.Println("Inside main")
	defer first()
	second()
	third()
}
func first(){
	fmt.Println("I will execute first")
}
func second(){
	fmt.Println("I will execute second")
}
func third(){
	fmt.Println("I will execute third")
}