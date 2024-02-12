package services

import (
	"crypto/tls"
	"fmt"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendMail(recipients []string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "mbaba317@gmail.com")
	m.SetHeader("To", recipients...)
	m.SetHeader("Subject", "Test")
	m.SetBody("text/plain", "This is the body plain text.")

	// m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	// m.Attach("/home/Alex/lolcat.jpg")

	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))

	fmt.Println(os.Getenv("SMTP_HOST"), port, os.Getenv("SMTP_USERNAME"), os.Getenv("SMTP_PASSWORD"))

	d := gomail.NewDialer(os.Getenv("SMTP_HOST"), port, os.Getenv("SMTP_USERNAME"), os.Getenv("SMTP_PASSWORD"))
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
