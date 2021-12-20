package liam

import (
	"bytes"
	"html/template"
)

// TODO: html template
type MailTemplate struct {
	mail     LMail
	template string
}

func (mt *MailTemplate) getMail() *LMail {
	return &mt.mail
}

func (mt *MailTemplate) Send() (err error) {
	return getResponse(mt)
}

func (mt *MailTemplate) getBodyBytes() (result []byte, err error) {
	header, err := mt.mail.bodyHeader()
	if err != nil {
		return nil, err
	}

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	header += mime

	t := template.Must(template.New("").Parse(mt.template))
	var body bytes.Buffer
	err = t.Execute(&body, mt.mail.body)
	if err == nil {
		body := body.Bytes()
		return append([]byte(header)[:], body[:]...), nil
	}
	return
}
