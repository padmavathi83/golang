package main

import "fmt"

func main(){
	sum(1,2)
	sum(1,3,6)

	nums := []int{3,4,6,7,8,9}
	sum(nums...)
}
func sum(nums ...int){
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums{
		total += num
	}
fmt.Println(total)
}

