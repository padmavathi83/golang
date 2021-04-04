package main

import "fmt"

func main(){
	defer print1("We will continue?")
	print2("we stop here")
}
func print1(s string){
	fmt.Println(s)
}
func print2(s string){
	fmt.Println(s)
}