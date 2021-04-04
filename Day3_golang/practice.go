package main

import (
	"fmt"
	"os"
)

func main() {

	if _, err := os.Stat("test.txt"); err == nil {
		fmt.Printf("File exists\n")
		fp.WriteString("\nWelcome to golang session from lkm")

		fmt.Printf("data append")

	} else {
		fmt.Printf("File does not exist\n")
		fp, _ := os.Create("test.txt")
		fmt.Println("File created", fp)
		fp.WriteString("\nWelcome to golang session from lkm")
		fmt.Println("File created")

	}
}
