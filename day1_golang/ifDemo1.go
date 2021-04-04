package main

import "fmt"

func main() {
	var num1,num2,num3 int
	fmt.Println("Enter 3 nos")
	fmt.Scanln(&num1,&num2,&num3)
	if num1 > num2 && num1 > num3{
		fmt.Println("num1 is bigger")
	}else if num2 > num1 && num2 > num3{
		fmt.Println("num2 is greater")
	}else {
		fmt.Println("num3 is greater")
	}
}