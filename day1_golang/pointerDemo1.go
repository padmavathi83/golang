package main

import ("fmt"
"reflect"
)

func main(){
	num1 := 10
	fmt.Println(num1)
	fmt.Println(&num1)
	//////////
	num2 := &num1
	fmt.Println(num2)
	////to know the type of num2
	fmt.Println(reflect.TypeOf(num2))
	//change value of num2
	*num2 = 200
	fmt.Println(num2)
	fmt.Println(*num2)
}