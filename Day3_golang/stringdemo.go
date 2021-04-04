package main

import ("fmt" 
"strings")

func main(){
	info := "My name is Rohit"
	fmt.Println(strings.Compare(info, "my name is Rohit"))
	fmt.Println(strings.Contains(info, "Rohit"))
	fmt.Println(strings.Contains(info, "Sony"))
	fmt.Println(strings.Count(info, "i"))
	fmt.Println(len(info))
	fmt.Println(strings.Index(info, "a"))  //starts from 0th 
	fmt.Println(strings.Split(info, " "))
	fmt.Println(strings.Split(info, ""))
	fmt.Println(strings.Split("hi,all,how,r;u?", ","))
	fmt.Println(strings.ToUpper(info))
	fmt.Println(strings.Title(info))
	fmt.Println(info)  //strings are immutable

	fmt.Println(strings.Replace("Hi new oldnew newold newstring","n","N",3))
	fmt.Println(strings.Repeat(info, 2))
	fmt.Println(info)
	//substring
	//a := info
	//b := a[2:5]
	fmt.Println(info[0:5])
}