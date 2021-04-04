package main

import "fmt"

func main(){
	fmt.Println(SaveDivide(10,2))
	fmt.Println(SaveDivide(10,0))
}
func SaveDivide(num1,num2 int) int{
	quotient := num1/num2
	return quotient
}