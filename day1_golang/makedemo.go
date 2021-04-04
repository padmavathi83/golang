package main

import "fmt"

func main(){
	myslice := make([]int,0)
/////
myslice = append(myslice, 22)
myslice = append(myslice, 23)
myslice = append(myslice, 24)
myslice = append(myslice, 25)
fmt.Println(myslice)
}