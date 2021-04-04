package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main(){
	wg.Add(2)  //adding and giving no of goroutines
	//wg will wait for 2 goroutines
	go sayHello()
	go sayHello()
	go sayBye()
	fmt.Println("waiting in main")
	wg.Wait()  //wait till goroutine acknowledge as Done
	fmt.Println("last line in main")
}
func sayHello(){
	for i:=0;i<5;i++{
		fmt.Println("Hello")
		time.Sleep(100*time.Millisecond)	
	}
	wg.Done()
}
func sayBye(){
	for i:=0;i<5;i++{
		fmt.Println("Bye")
		time.Sleep(100*time.Millisecond)	
	}
	wg.Done()
}