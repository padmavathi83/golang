package main

import "fmt"

func NewSlice(start, count, step int)[]int {
s := make([]int, count)
for i:= range s{
	s[i] = start
	start += step
}
return s
}
func main(){
s := []int{0,1,2,3,4,5,6,7,8,9}
fmt.Println(s)
/////
fmt.Println(NewSlice(10,10,1))
fmt.Println(NewSlice(10,10,2))
fmt.Println(NewSlice(1,10,3))
fmt.Println(NewSlice(-1,10,-1))
fmt.Println(NewSlice(10,10,0))
}