package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func main() {
	// Configuration
	//from := "porehg.93854@gmail.com"
	//password := "9IyVplQ9cLZN"
	//to := []string{"x-bit89@mail.ru"}
	//smtpHost := "smtp.gmail.com"
	//smtpPort := "587"
	//ds
	//message := []byte("My super secret message.")

	// Create authentication
	//auth := smtp.PlainAuth("", "porehg.93854@gmail.com", "9IyVplQ9cLZN", "smtp.gmail.com")
	//// Send actual message
	msg := []byte("Subject: " + "Restore password" + "\r\n\r\n" + "sdfsdfsdfsdfsdfsdfsdf: " + "http://10.10.15.185:5000/internal/swagger" + "?" + "sdfsf")
	////err := smtp.SendMail("smtp.gmail.com"+":"+"587", auth, "porehg.93854@gmail.com", []string{"x-bit89@mail.ru"}, []byte("My super secret message."))
	//err := smtp.SendMail("smtp.gmail.com"+":"+"587", auth, "porehg.93854@gmail.com", []string{"x-bit89@mail.ru"}, msg)
	//
	//if err != nil {
	//	fmt.Println("Error not nil")
	//	log.Fatal(err)
	//}

	auth2 := smtp.PlainAuth("", "porehg.93854@gmail.com", "9IyVplQ9cLZN", "smtp.gmail.com")
	err := smtp.SendMail("smtp.gmail.com"+":"+"587", auth2, "porehg.93854@gmail.com", []string{"sorovskii84@gmail.com"}, []byte(msg))
	if err != nil {
		fmt.Println("Error not nil")
		log.Fatal(err)
	}

	//msg := fmt.Sprintf("%s: %s\n\n%s: %s?%s", "Subject", "Password restore", "Link to change password", "http://10.10.15.185:5000/internal/swagger", "sdfsf")
	//auth2 := smtp.PlainAuth("", "x-bit89@mail.ru", "marseek6726984", "smtp.mail.ru")
	//fmt.Println("start send")
	//err := smtp.SendMail("smtp.mail.ru"+":"+"465", auth2, "x-bit89@mail.ru", []string{"sorovskii84@gmail.com"}, []byte(msg))
	//if err != nil {
	//	fmt.Println("Error not nil")
	//	log.Fatal(err)
	//}
	//fmt.Println("end send")
}
