package main

import ("fmt"
"time"
)

func main(){
	ptime := time.Now()
	fmt.Println(ptime)

	fmt.Println(ptime.Date())  //only date

	fmt.Println(ptime.Day())
	fmt.Println(ptime.Hour())
	fmt.Println(ptime.Minute())
	fmt.Println(ptime.Local())
	time.Sleep(5000*time.Microsecond)
	fmt.Println(ptime.Clock())

	wday := time.Now().Weekday()
	fmt.Println(wday)
	fmt.Println(int(wday))

}