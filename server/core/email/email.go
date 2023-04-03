package email

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/global"
	"html/template"
	"io/ioutil"
	"net/smtp"
	"os"
	"strings"
)

type Interface interface {
	Send(email, subject, context string) error
	SendAll(emails []string, subject, context string) error
}

type email struct {
	template map[string]string
}

type sender struct {
	template string
	user     string
	host     string
	password string
	title    string
}

var (
	Templates = make(map[string]string)
)

const (
	defaultTemplate = "default"
)

func Init() {
	for name, fileName := range global.Config.Email.Template {
		file, err := os.Open(fileName)
		if err != nil {
			panic("邮箱模板初始化失败:" + err.Error())
		}
		key, err := ioutil.ReadAll(file)
		if err != nil {
			panic("读取邮箱模板失败:" + err.Error())
		}
		Templates[name] = string(key)
	}
}

func New(tp ...string) Interface {
	t := ""
	if len(tp) != 0 && tp[0] != "" {
		t = Templates[tp[0]]
	} else {
		t = Templates[defaultTemplate]
	}

	return &sender{
		user:     global.Config.Email.User,
		host:     global.Config.Email.Host,
		password: global.Config.Email.Password,
		title:    global.Config.Email.Company,
		template: t,
	}
}

func (e *sender) parseTemplate(context string) string {
	n := template.New("")
	t, err := n.Parse(e.template)
	if err != nil {
		return context
	}

	html := bytes.NewBuffer([]byte(""))
	if err = t.Execute(html, gin.H{
		"content": context,
	}); err != nil {
		return context
	}
	return html.String()
}

func (e *sender) Send(email string, subject, context string) error {
	context = e.parseTemplate(context)
	hp := strings.Split(e.host, ":")
	auth := smtp.PlainAuth("", e.user, e.password, hp[0])
	ct := "Content-Type: text/html; charset=UTF-8"
	msg := []byte("To: " + email + "\r\nFrom: " + e.title + "\r\nSubject: " + subject + "\r\n" + ct + "\r\n\r\n" + context)
	return smtp.SendMail(e.host, auth, e.user, []string{email}, msg)
}

func (e *sender) SendAll(emails []string, subject, context string) error {
	context = e.parseTemplate(context)
	hp := strings.Split(e.host, ":")
	auth := smtp.PlainAuth("", e.user, e.password, hp[0])
	ct := "Content-Type: text/html; charset=UTF-8"
	to := strings.Join(emails, ";")
	msg := []byte("To: " + to + "\r\nFrom: " + e.title + ">\r\nSubject: " + subject + "\r\n" + ct + "\r\n\r\n" + context)
	return smtp.SendMail(e.host, auth, e.user, emails, msg)
}
