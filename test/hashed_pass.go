package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main6() {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password1234"), 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(hashedPassword))

}
