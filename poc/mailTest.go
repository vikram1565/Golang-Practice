package main

import (
	"crypto/tls"

	mail "gopkg.in/gomail.v2"
)

func main() {
	m := mail.NewMessage()
	m.SetHeader("From", "akshayb@mkcl.org")
	m.SetHeader("To", "vikrami@mkcl.org")
	m.SetHeader("Reply-To", "suvarnar@mkcl.org")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/plain", "This is the body of the message.")

	d := mail.Dialer{Host: "10.1.70.100", Port: 25, Username: "", Password: "", SSL: false}
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
