package main

import "fmt"

func main(){
add(3,4)
mul, div := multi(12,3)
fmt.Println("Multiplication of 2 nos is",mul)
fmt.Println("Division of two nos is",div)
}
func add(num1 int,num2 int){
	fmt.Println("sum is",num1+num2)
}
func multi(n1 int,n2 int) (int,int){
return n1 * n2, n1/n2
}