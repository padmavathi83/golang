package main

import "fmt"

func main(){
mymap := map[int]string{101:"Sam",102:"Ram"}
fmt.Println(mymap)
///
mymap[444] = "Paul"
fmt.Println(mymap)
mymap[102] = "Krish"
fmt.Println(mymap)

fmt.Println(mymap[101])

//delete
delete(mymap,444)
fmt.Println(mymap)
}