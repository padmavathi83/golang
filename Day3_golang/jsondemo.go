package main

import (
	"encoding/json"
	"fmt"
)

func main(){
	text := "[{\"id\":100,\"name\":\"Paul\"}]"
	jsondata := []byte(text)
	var p[]person
	json.Unmarshal(jsondata, &p)
	fmt.Println(p)
}
type person struct{
	Id int
	Name string
}