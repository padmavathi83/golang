package main

import ("fmt"
"runtime"
)

func main(){
	fmt.Println(runtime.GOOS)  //os details
	fmt.Println(runtime.NumCPU())  //gives number of core
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(runtime.GOROOT())
	fmt.Println(runtime.Compiler)
}