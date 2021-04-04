package main

import "fmt"

type address struct{
	pin int
	city string
}
type student struct{
	id int
	name string
	age int
	sadd address	
}

func main(){
	//create address struct and pass it to student
	//addobj := address{570032,"Mysore"}
	//pass it to address type in student
	//sobj := student{101, "Scooby",20,addobj}
	//fmt.Println(sobj)
	//fmt.Println(sobj.id,sobj.name,sobj.sadd.city)
	//keyvalue format
	s1obj := student{id:200,age:18,name:"Rohit",sadd:address{560012,"Bangalore"}}
	fmt.Println(s1obj)
}