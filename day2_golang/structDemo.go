package main 

import "fmt" 

type employee struct { 
firstName string 
lastName string 
age int 
} 
func main() { 
x := employee{age: 30, firstName: "John", lastName: "Anderson"} 
//fmt.Println(x) 
fmt.Println(x.firstName)
fmt.Println(x.lastName)
fmt.Println(x.age)
} 
