package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, _ := os.Create("Data.txt")
	fmt.Println("File created", f)
	bdata := []byte("\nwelcome to GoLang session\n")
	//now it is in write mode
	f.Write(bdata)
	fmt.Println("data written")
	//another way to write data
	f.WriteString("\n Session location is in BDC")
	fmt.Println("data appended")
	f.Close() //close file

	//read data
	data, _ := ioutil.ReadFile("Data.txt")
	fmt.Println(string(data))

	//condition
	fp, _ := os.OpenFile("example.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	fp.WriteString("\n Golang is used in Google Cloud")
}
