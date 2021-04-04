package main

import (
	"fmt"
	"reflect"
)

var name = "Rohit" //initialization and assign value
var num float64    //initialization
var num1 int

func main() {
	flag := true
	age := 20 //local
	fmt.Println("Hi all welcome to GoLang" + " " + name)
	num = 10.2
	fmt.Println("The num value is", num)
	fmt.Println(age)
	//to know the data type of variable
	fmt.Print(reflect.TypeOf(age))
	fmt.Print(reflect.TypeOf(flag))
	fmt.Printf("\n%T", age)
	fmt.Printf("\n%T", name)
	fmt.Printf("\n%T", num)
}
