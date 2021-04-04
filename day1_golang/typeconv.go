package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 42
	fmt.Printf("%v, %T\n", i, i)

	var j string
	j = string(i)  //when u run this code it will print *
	j = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", j, j)
}
