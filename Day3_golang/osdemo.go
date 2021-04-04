package main

import ("fmt"
"os"
)

func main(){
	fp, _ := os.Create("C:/Users/nagalatha.b.raj/Desktop/newfile.txt")
	fmt.Println("File created", fp)
	// create folder
	fd := os.Mkdir("Sample",0644)  //folder permission
	fmt.Println("dir created", fd)

}