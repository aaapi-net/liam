package liam

import "fmt"

type MailText struct {
	mail LMail
}

func (ms *MailText) getMail() *LMail {
	return &ms.mail
}

func (ms *MailText) getBodyBytes() ([]byte, error) {
	return []byte(fmt.Sprintf("%+v", ms.mail.body)), nil
}

func (ms *MailText) Send() (err error) {
	return getResponse(ms)
}
