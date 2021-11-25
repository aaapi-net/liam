package liam

import (
	"bytes"
	"html/template"
)
// TODO: html template
type MailTemplate struct {
	mail LMail
	template string
}

func (mt *MailTemplate) getMail() *LMail {
	return &mt.mail
}

func (mt *MailTemplate) Send() (err error) {
	return getResponse(mt)
}

func (mt *MailTemplate) getBodyBytes() (result []byte, err error) {
	t := template.Must(template.New("").Parse(mt.template))
	var body bytes.Buffer
	err = t.Execute(&body, mt.mail.body)
	if err == nil {
		return body.Bytes(), nil
	}
	return
}