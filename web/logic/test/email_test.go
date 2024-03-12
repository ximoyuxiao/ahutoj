package logic

import (
	"ahutoj/web/io/request"
	"fmt"
	"html/template"
	"testing"
	// "math/rand"
	// "time"
	"github.com/wneessen/go-mail"
)
const (
		htmlBodyTemplate = 
	`<p>Hi {{.Uname}},</p>
	<p>Welcome to AhutOJ!</p>
	<p>Please confirm your email address by clicking on the button below.</p>
	<a href="{{.Link}}" style="
  font-size:16px;
  font-weight:700;
  padding:15px 40px;
  color:#fff;
  background-color:#2595ff;
  border-color:#0b89ff;
  text-decoration:none;
  display:inline-block;
  margin-bottom: 0;
  text-align: center;
  vertical-align: middle;
  touch-action: manipulation;
  cursor: pointer;
  background-image: none;
  border: 1px solid transparent;
  white-space: nowrap;
  line-height: 1.428571429;
  border-radius: 4px;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;" rel="noopener" target="_blank">Confirm your email</a>
	<p>Happy coding!</p>
	<p>By AhutACM.</p>`
)
var user =request.VerifyEmailReq{
	Email: "",
	UID:  "22904301",
	Uname: "ahutoj",
}
var ServerName="AhutOJ"
var ServerAddr="ahutoj-verify@jorban.top"
// var ServerAddr2="2648242688@qq.com"
var PassWord="HzSMXKkkB2PEfMDf"
// var PassWord2="ccbaafuwcqjwecdi"
type TemplateData struct {
	Uname string
	Link string
}
func TestEmailSend(t *testing.T) {
	//模板注入参数
	data:=TemplateData{
		Uname: user.Uname,
		Link:  "http://localhost:8080/verify?uid="+user.UID,
	}
	//邮件模板
	htp, err := template.New("htp").Parse(htmlBodyTemplate)//使用template注入更加安全
	if err != nil {
		t.Fatalf("failed to parse text template: %s", err)
	}
	//邮件格式
	email := mail.NewMsg()
	if err := email.FromFormat(ServerName,ServerAddr); err != nil {
		t.Fatalf("failed to set formatted FROM address: %s", err)
	}
	if err := email.To(user.Email); err != nil {
		t.Fatalf("failed to set To address: %s", err)
	}
	//内容
	email.Subject(fmt.Sprintf("Hi %s ,Welcome to AhutOJ! Please confirm your email address", user.Uname))
	if err := email.SetBodyHTMLTemplate(htp,data); err != nil {
		t.Fatalf("failed to set HTML template as HTML body: %s", err)
	}
	c, err := mail.NewClient("smtp.exmail.qq.com",
		mail.WithSMTPAuth(mail.SMTPAuthPlain), mail.WithSSLPort(true),
		mail.WithUsername(ServerAddr), mail.WithPassword(PassWord),
	)
	if err != nil {
		t.Fatalf("failed to create client: %s", err)
	}
	defer c.Close()
	if err := c.DialAndSend(email); err != nil {
		t.Fatalf("failed to deliver mail: %s", err)
	}
}
