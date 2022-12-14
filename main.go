package main

import (
	"fmt"
	"net/smtp"

	"github.com/aziemp66/go-mail/helper"
)

func sendMailSimple(mail string, appPassword string, to []string, subject string, body string) {
	auth := smtp.PlainAuth("", mail, appPassword, "smtp.gmail.com")

	msg := []byte("Subject: " + subject + "\r" + "\r" + body + "\r")

	err := smtp.SendMail("smtp.gmail.com:587", auth, mail, to, msg)

	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func main() {
	cfg := helper.LoadConfig()

	sendMailSimple(cfg.APP_EMAIL, cfg.APP_PASSWORD, []string{"azielala55@gmail.com"}, "Hello", "Welcome to SMTP With Go")
}
