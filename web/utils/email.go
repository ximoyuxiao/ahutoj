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

func SendEmail(r *Receiver,data *TemplateData,htp *template.Template)error {
	//邮件格式
	var emcfg=GetConfInstance().Email
	// logger.Debugf("%+v",Sdump(emcfg))
	Email := mail.NewMsg()
	if err := Email.From(emcfg.Serveremail); err != nil {
		logger.Errorf("failed to set formatted FROM address: %s", err)
		return err
	}
	if err := Email.To(r.Email); err != nil {
		logger.Errorf("failed to set To address: %s", err)
		return err
	}
	//内容
	Email.Subject(fmt.Sprintf("Hi %s ,Welcome to AhutOJ! Please confirm your email address", r.Username))
	if err := Email.SetBodyHTMLTemplate(htp,data); err != nil {
		logger.Errorf("failed to set HTML template as HTML body: %s", err)
		return err
	}
	c, err := mail.NewClient(emcfg.Stmpaddr,
		mail.WithSMTPAuth(mail.SMTPAuthPlain), mail.WithSSLPort(true),
		mail.WithUsername(emcfg.Serveremail), mail.WithPassword(emcfg.Password),
	)
	if err != nil {
		logger.Errorf("failed to create client: %s", err)
		return err
	}
	defer c.Close()
	// err = c.DialAndSend(Email)
	// if err != nil {
	// 	logger.Errorf("failed to deliver mail: %s", err)
	// 	return err
	// }
	return nil
}
func EmailVerify(uname string,token string,email string,path string)error{
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
		return err
	}
	err =SendEmail(&r,&data,htp)
	if err!=nil{
		logger.Errorf("failed to send email: %s", err)
		return err
	}
	return nil
}