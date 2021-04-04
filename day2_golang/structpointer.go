package main

import "fmt"

type emp struct{
id int
name string
age float32
}

func main(){
fmt.Println(emp{name:"Rohit",age:18,id:1234})

e := emp{name:"Ram",age:28,id:1236}
fmt.Println(e.name)

//assigning pointer to structure
emp_ptr := &e
fmt.Println(emp_ptr.id)

//update name
emp_ptr.age = 29
fmt.Println(emp_ptr.age)
}