package main

import "fmt"

func main(){
marksgot := studmarks(89, 76, 79,97,70,80)
fmt.Println(marksgot)
}

func studmarks(marks ...int) int{
m := marks[0]
for _, i := range marks{
	if i < m{
		m = i
	}
}
return m
}
