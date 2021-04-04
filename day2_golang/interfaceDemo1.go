package main

import "fmt"

//create interface
type vehicle interface {
	engine()
}

//create function
func vdetails(v vehicle) {
	fmt.Println(v)
	//v.engine()
}

//create structue
type audi struct {
	model string
	name  string
}
type maruthi struct {
	model string
	name  string
}

//override function then it will become vehicle
func (a audi) engine() {
	fmt.Println("Audi had auto engine")
}
func (m maruthi) engine() {
	fmt.Println("Maruthi has manual/auto engine")
}
func main() {
	aobj := audi{"B110", "Audi New"}
	aobj.engine()
	mobj := maruthi{"M110", "DZire"}
	vdetails(mobj)
}
