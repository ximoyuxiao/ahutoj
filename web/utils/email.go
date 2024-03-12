package utils

import (
	"fmt"
	"html/template"
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
type TemplateData struct {
	Uname string
	Link string
}
type Receiver struct {
	Email string
	Username string
}
var emcfg=GetConfInstance().Email

func SendEmail(r *Receiver,data *TemplateData,htp *template.Template) {
	//邮件格式
	Email := mail.NewMsg()
	if err := Email.From(emcfg.Serveremail); err != nil {
		logger.Errorf("failed to set formatted FROM address: %s", err)
	}
	if err := Email.To(r.Email); err != nil {
		logger.Errorf("failed to set To address: %s", err)
	}
	//内容
	Email.Subject(fmt.Sprintf("Hi %s ,Welcome to AhutOJ! Please confirm your email address", r.Username))
	if err := Email.SetBodyHTMLTemplate(htp,data); err != nil {
		logger.Errorf("failed to set HTML template as HTML body: %s", err)
	}
	c, err := mail.NewClient(emcfg.Stmpaddr,
		mail.WithSMTPAuth(mail.SMTPAuthPlain), mail.WithSSLPort(true),
		mail.WithUsername(emcfg.Serveremail), mail.WithPassword(emcfg.Password),
	)
	if err != nil {
		logger.Errorf("failed to create client: %s", err)
	}
	defer c.Close()
	if err := c.DialAndSend(Email); err != nil {
		logger.Errorf("failed to deliver mail: %s", err)
	}
}
func EmailVerify(uname string,email string,token string,path string){
	data:=TemplateData{
		Uname: uname,
		Link: path+"?email="+email+"&token="+token,
	}
	r:=Receiver{
		Email: email,
		Username: uname,
	}
	//邮件模板
	htp, err := template.New("htp").Parse(htmlBodyTemplate)//使用template注入更加安全
	if err != nil {
		logger.Errorf("failed to parse text template: %s", err)
	}
	SendEmail(&r,&data,htp)
}