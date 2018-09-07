package models

import (
	"time"

	"github.com/astaxie/beego"
	gomail "gopkg.in/gomail.v2"
)

//Email ...
type Email struct {
	To          []string
	Bcc         []string
	Cc          []string
	Subject     string
	HTML        string
	Headers     [][]string
	Attachments []string
}

var (
	from    = beego.AppConfig.String("email::from")
	name    = beego.AppConfig.String("email::name")
	pass    = beego.AppConfig.String("email::pass")
	host    = beego.AppConfig.String("email::host")
	port, _ = beego.AppConfig.Int("email::port")
)

//SendMails ...
func SendMails(Emails []*Email) {

	ch := make(chan *gomail.Message)

	for _, Email := range Emails {
		go SendMail(Email)
	}

	close(ch)
}

//SendMail ...
func SendMail(Email *Email) error {
	m := gomail.NewMessage()

	for _, header := range Email.Headers {
		m.SetHeader(header[0], header[1])
	}

	for _, Attachment := range Email.Attachments {
		m.Attach(Attachment)
	}

	m.SetHeaders(map[string][]string{
		"From":    {m.FormatAddress(from, name)},
		"To":      Email.To,
		"Cc":      Email.Cc,
		"Bcc":     Email.Bcc,
		"Subject": {Email.Subject},
		"X-Date":  {m.FormatDate(time.Now())},
	})

	m.SetBody("text/html", Email.HTML)

	d := gomail.NewDialer(host, port, from, pass)

	err := d.DialAndSend(m)

	if err != nil {
		return err
	}

	return nil

}
