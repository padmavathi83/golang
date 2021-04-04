package main

import "fmt"

type Shape interface{
	area() float64
}

type Rectangle struct{
	width,height float64
}
func (r Rectangle) area() float64{
	return r.width * r.height
}
func getArea(s Shape) float64{
	return s.area()
}
func main(){
	rect := Rectangle{width:10,height: 5}
	fmt.Printf("Area of Rectangle: %f\n", getArea(rect))
}