package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"

	"github.com/aziemp66/go-mail/helper"
	"github.com/aziemp66/go-mail/mail"
	"gopkg.in/gomail.v2"
)

func sendMailSimple(mail string, appPassword string, to []string, subject string, body string) {
	auth := smtp.PlainAuth("", mail, appPassword, "smtp.gmail.com")

	msg := []byte("Subject: " + subject + "\r" + "\r" + body + "\r")

	err := smtp.SendMail("smtp.gmail.com:587", auth, mail, to, msg)

	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func sendMailSimpleHtml(mail string, appPassword string, to []string, subject string, templatePath string) {
	//get html
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)

	t.Execute(&body, map[string]string{
		"Name": "Azie",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

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

	message += "\r" + body.String()

	err = smtp.SendMail("smtp.gmail.com:587", auth, mail, to, []byte(message))

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

}

func sendGoMail(mail string, appPassword string, to []string, subject string, templatePath string, attachmentPath string) {
	//get html
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)

	t.Execute(&body, map[string]string{
		"Name": "Azie",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	//send with gomail
	m := gomail.NewMessage()
	m.SetHeader("From", mail)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body.String())
	m.Attach(attachmentPath)

	d := gomail.NewDialer("smtp.gmail.com", 587, mail, appPassword)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

func main() {
	cfg := helper.LoadConfig()

	// sendMailSimple(cfg.APP_EMAIL, cfg.APP_PASSWORD, []string{"azielala55@gmail.com"}, "Hello", "Welcome to SMTP With Go")
	// sendMailSimpleHtml(cfg.APP_EMAIL, cfg.APP_PASSWORD, []string{"aziemp55@gmail.com"}, "Hello", "./template/invoice.gohtml")
	// sendGoMail(cfg.APP_EMAIL, cfg.APP_PASSWORD, []string{"azielala55@gmail.com"}, "GoMail, The Third Party Mail Server", "./template/invoice.gohtml", "./picture/92568390.jpg")
	//get html
	var body bytes.Buffer
	t, err := template.ParseFiles("./template/invoice.gohtml")

	if err != nil {
		panic(err)
	}

	t.Execute(&body, map[string]string{
		"Name": "Azie",
	})

	sender := mail.New(cfg.APP_EMAIL, cfg.APP_PASSWORD, "smtp.gmail.com", "587")
	m := mail.NewMessage("Test Go Mail", body.String())
	m.To = []string{"aziemp55@gmail.com"}
	m.AttachFile("./picture/92568390.jpg")

	fmt.Println(sender.Send(m))
}
