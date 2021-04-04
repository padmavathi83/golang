package main

import (
	"fmt"
	"time"
)


func main(){
	 say("Hello")
	 go say("hi")
}

func say(s string){
	for i:=0;i<5;i++{
		time.Sleep(100* time.Millisecond)
		fmt.Println(s)
	}
}