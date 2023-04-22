package email

import (
	"bytes"
	"github.com/limeschool/easy-admin/server/config"
	"github.com/limeschool/easy-admin/server/errors"
	"html/template"
	"io/ioutil"
	"net/smtp"
	"os"
	"strings"
)

type email struct {
	template map[string]struct {
		subject string
		html    string
	}
	user     string
	host     string
	password string
	company  string
}

type Email interface {
	NewSender(tpName string) Sender
}

type sender struct {
	tp string
	*email
}

type Sender interface {
	Send(email string, data any) error
	SendAll(emails []string, data any) error
}

// New 初始化email实例
func New(conf config.Email) Email {
	emailIns := email{
		user:     conf.User,
		host:     conf.Host,
		password: conf.Password,
		company:  conf.Company,
		template: map[string]struct {
			subject string
			html    string
		}{},
	}

	for _, item := range conf.Template {
		file, err := os.Open(item.Src)
		if err != nil {
			panic("邮箱模板初始化失败:" + err.Error())
		}
		val, err := ioutil.ReadAll(file)
		if err != nil {
			panic("读取邮箱模板失败:" + err.Error())
		}
		emailIns.template[item.Name] = struct {
			subject string
			html    string
		}{subject: item.Subject, html: string(val)}
	}

	return &emailIns
}

// NewSender
//
//	@Description: 新建一个发送器
//	@receiver e
//	@param tpName 需要发送的模板名
//	@return Sender
func (e *email) NewSender(tpName string) Sender {
	return &sender{
		email: e,
		tp:    tpName,
	}
}

// Send
//
//	@Description: 向指定的邮箱发送邮件
//	@receiver e
//	@param email 需要发送的邮箱
//	@param data 模板参数
//	@return error
func (e *sender) Send(email string, data any) error {
	if !e.hasTemplate() {
		return errors.New("not exist template")
	}
	html, err := e.parseTemplate(data)
	if err != nil {
		return err
	}
	hp := strings.Split(e.host, ":")
	auth := smtp.PlainAuth("", e.user, e.password, hp[0])
	ct := "Content-Type: text/html; charset=UTF-8"
	subject := e.email.template[e.tp].subject
	msg := []byte("To: " + email + "\r\nFrom: " + e.company + "\r\nSubject: " + subject + "\r\n" + ct + "\r\n\r\n" + html)
	return smtp.SendMail(e.host, auth, e.user, []string{email}, msg)
}

// SendAll
//
//	@Description: 批量发送邮件信息
//	@receiver e
//	@param emails 需要发送的邮箱列表
//	@param data  模板填充变量
//	@return error
func (e *sender) SendAll(emails []string, data any) error {
	if !e.hasTemplate() {
		return errors.New("not exist template")
	}
	html, err := e.parseTemplate(data)
	if err != nil {
		return err
	}
	hp := strings.Split(e.host, ":")
	auth := smtp.PlainAuth("", e.user, e.password, hp[0])
	ct := "Content-Type: text/html; charset=UTF-8"
	to := strings.Join(emails, ";")
	subject := e.email.template[e.tp].subject
	msg := []byte("To: " + to + "\r\nFrom: " + e.company + ">\r\nSubject: " + subject + "\r\n" + ct + "\r\n\r\n" + html)
	return smtp.SendMail(e.host, auth, e.user, emails, msg)
}

// hasTemplate
//
//	@Description: 判断是否存在模板
//	@receiver e
//	@return bool
func (e *sender) hasTemplate() bool {
	_, is := e.email.template[e.tp]
	return is
}

// parseTemplate
//
//	@Description: 解析模板变量
//	@receiver e
//	@param data
//	@return string
func (e *sender) parseTemplate(data any) (string, error) {
	n := template.New("")
	htmlTemplate := e.email.template[e.tp].html
	t, err := n.Parse(htmlTemplate)
	if err != nil {
		return "", err
	}
	html := bytes.NewBuffer([]byte(""))
	if err = t.Execute(html, data); err != nil {
		return "", err
	}
	return html.String(), nil
}
