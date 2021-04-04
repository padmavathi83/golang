package main

import "fmt"

func main() {
ch := "i"
switch ch{
case "a":fmt.Println("Vowel", ch)
case "e":fmt.Println("Vowel", ch)
case "i":fmt.Println("Vowel", ch)
case "o":fmt.Println("Vowel", ch)
case "u":fmt.Println("Vowel", ch)
default:fmt.Println("Not a Vowel")
}
}