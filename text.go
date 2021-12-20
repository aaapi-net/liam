package liam

import "fmt"

type MailText struct {
	mail LMail
}

func (ms *MailText) getMail() *LMail {
	return &ms.mail
}

func (ms *MailText) getBodyBytes() ([]byte, error) {
	header, err := ms.mail.bodyHeader()
	if err != nil {
		return nil, err
	}

	body := []byte(fmt.Sprintf("%+v", ms.mail.body))

	return append([]byte(header)[:], body[:]...), nil
}

func (ms *MailText) Send() (err error) {
	return getResponse(ms)
}
