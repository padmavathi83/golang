package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	pass := "Admin"
	newpass, _ := bcrypt.GenerateFromPassword([]byte(pass), 4) //encrypt Admin password
	//4 times wrap it
	fmt.Println(newpass)
}
