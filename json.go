package liam

import (
	"encoding/json"
)

type MailJson struct {
	mail LMail
}

func (mj *MailJson) getMail() *LMail {
	return &mj.mail
}

func (mj *MailJson) getBodyBytes() ([]byte, error) {
	header, err := mj.mail.bodyHeader()
	if err != nil {
		return nil, err
	}

	body, err := json.Marshal(mj.mail.body)

	return append([]byte(header)[:], body[:]...), err
}

func (mj *MailJson) Send() (err error) {
	return getResponse(mj)
}
