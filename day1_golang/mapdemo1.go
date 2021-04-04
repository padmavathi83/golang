package main

import "fmt"

func main(){
	//map value can also be initialized using make function
	var employee = make(map[string]int)
	employee["John"] = 10
	employee["Rohit"] = 20
	employee["Pinky"] = 30
	fmt.Println(employee)

	//

	employeeList := make(map[string]int)
	employeeList["John"] = 10
	employeeList["Rohit"] = 20
	fmt.Println(employeeList)
}