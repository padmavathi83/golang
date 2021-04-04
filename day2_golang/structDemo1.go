package main

import "fmt"

type emp struct{
id int
name string
age float32
}

func main(){
	eobj := emp{}
	eobj.id = 1234
	eobj.name = "Rohit"
	eobj.age = 22.5
	fmt.Println(eobj)
	///////////
	//pass all values
	eobj1 := emp{1235, "Pinky", 18.5}
	fmt.Println(eobj1)
	///// assign values using key value pair
	eobj2 := emp{id:1236}
	fmt.Println(eobj2)
	//
	eobj.details()
}
//creating function with receiver
func (eobj emp)details(){
	fmt.Println("Employee details are ",eobj.id," ",eobj.name," ",eobj.age)
}