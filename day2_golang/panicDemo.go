package main

import "os"

func main(){
	panic("error situation")
	_, err := os.Open("/tmp/file")
	if err != nil{
		panic(err)
	}
}