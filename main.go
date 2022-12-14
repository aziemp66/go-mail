package main

import (
	"fmt"
	"net/smtp"

	"github.com/aziemp66/go-mail/helper"
)

func sendMailSimple(mail string, appPassword string, to string, subject string, body string) {
	auth := smtp.PlainAuth("", mail, appPassword, "smtp.gmail.com")
	toList := []string{to}

	msg := []byte("To: " + to + "\r" + "Subject: " + subject + "\r" + "\r" + body + "\r")

	err := smtp.SendMail("smtp.gmail.com:587", auth, mail, toList, msg)

	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func main() {
	cfg := helper.LoadConfig()

	sendMailSimple(cfg.APP_EMAIL, cfg.APP_PASSWORD, "azielala55@gmail.com", "Test", "Welcome to Go")
}
