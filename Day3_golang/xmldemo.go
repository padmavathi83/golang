package main

import ("fmt"
"encoding/xml"
)
//xml data
var myxml = []byte(
	`<emp>
	<fname>Raj</fname>
	<age>35</age>
	<loc>BDC</loc>
	<mob>9876456790</mob>
	</emp>`)
//create structure
type employee struct{
	//all variables should be public so use first letter capital
	Ename string `xml:"fname"`
	Eage int `xml:"age"`
	Eloc string `xml:"loc"`
	Emob int `xml:"mob"`
}
func main(){
	var eobj employee
	xml.Unmarshal(myxml, &eobj)
	fmt.Println(eobj)
}