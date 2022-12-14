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

func sendMailSimpleHtml(mail string, appPassword string, to []string, subject string, html string) {
	auth := smtp.PlainAuth("", mail, appPassword, "smtp.gmail.com")

	header := make(map[string]string)
	header["From"] = mail
	header["To"] = to[0]
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/html; charset=\"UTF-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r", k, v)
	}

	message += "\r" + html

	err := smtp.SendMail("smtp.gmail.com:587", auth, mail, to, []byte(message))

	if err != nil {
		fmt.Println("Error: ", err)
	}

}

func main() {
	cfg := helper.LoadConfig()

	// sendMailSimple(cfg.APP_EMAIL, cfg.APP_PASSWORD, []string{"azielala55@gmail.com"}, "Hello", "Welcome to SMTP With Go")
	sendMailSimpleHtml(cfg.APP_EMAIL, cfg.APP_PASSWORD, []string{"aziemp55@gmail.com"}, "Hello", "<h1>Welcome to SMTP With Go</h1><p>My Name Is Gustavo Fring, You Can Call me Gus</p>")
}
