package main

import "fmt"

func display(done chan bool){
	fmt.Println("Display goroutine")
	done <- true
}
func main(){
	done := make(chan bool)
	go display(done)
	<- done
	fmt.Println("main function")
}