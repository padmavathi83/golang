package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, _ := os.Create("Data.txt")
	fmt.Println("file created", f)
	bdata := []byte("hello paddu. how are you \n")
	f.Write(bdata)
	fmt.Println("data written")
	f.WriteString("/n , I am fine")

	f.Close()
	data, _ := ioutil.ReadFile("Data.txt")
	fmt.Println(string(data))

	fp, _ := os.OpenFile("sample.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	fp.WriteString("hello Teja")

}
