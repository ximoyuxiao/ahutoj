package utils

import (
	"fmt"
	"net/smtp"
	"strings"
)

var SmtpUser struct {
	auth     smtp.Auth
	user     string
	password string
	Server   string
	Port     string
}

func LoginSmtp(user string, password string, smtpServer string, smtpPort string) {
	SmtpUser.auth = smtp.PlainAuth("", user, password, smtpServer)
	SmtpUser.user = user
	SmtpUser.password = password
	SmtpUser.Server = smtpServer
	SmtpUser.Port = smtpPort
}
func SendEmail(touser []string, title string, text string) error {
	message := ""
	message += fmt.Sprintf("To: %s\r\n", strings.Join(touser, ","))
	message += fmt.Sprintf("Subject: %s\r\n\r\n", title)
	message += text
	return smtp.SendMail(SmtpUser.Server+":"+SmtpUser.Port, SmtpUser.auth, SmtpUser.user, touser, []byte(message))
}
