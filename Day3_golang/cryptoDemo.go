package main

import ("fmt"
"golang.org/x/crypto/bcrypt"
)

func main(){
	pass := "Admin"
	newpass, _ := bcrypt.GenerateFromPassword([]byte(pass),4)  //encrypt Admin password
	//4 times wrap it
	fmt.Println(newpass)
	//decryption
	err := bcrypt.CompareHashAndPassword(newpass,[]byte("admin"))
	//if both matches then ok or error
	if err != nil {
		fmt.Println("Invalid password")
	}else{
		fmt.Println("welcome user password matched login now")
	}
}