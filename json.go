package liam

import (
	"crypto/tls"
	"encoding/json"
	"log"
	"net/smtp"
)

type MailJson struct {
	mail LMail
}

func (mj *MailJson) getMail() *LMail {
	return &mj.mail
}

func (mj *MailJson) getBodyBytes() ([]byte, error) {
	return json.Marshal(mj.mail.body)
}

func (mj *MailJson) Send() (err error) {
	return getResponse(mj)
}

func (mj *MailJson) Response2() (err error) {
	header, err := mj.mail.bodyHeader()

	if err == nil {
		body, err := json.Marshal(mj.mail.body)
		auth := smtp.PlainAuth("", mj.mail.smtp.username, mj.mail.smtp.password, mj.mail.smtp.server)

		log.Println("Header: ", header)

		if err == nil {
			conn, err := tls.Dial("tcp", mj.mail.smtp.getAddr(), mj.mail.smtp.liamog.tlsConfig)
			if err != nil {
				return err
			}

			client, err := smtp.NewClient(conn, mj.mail.smtp.server)

			if err != nil {
				return err
			}

			if err = client.Auth(auth); err != nil {
				return err
			}

			writer, err := client.Data()
			if err != nil {
				return err
			}
			defer writer.Close()

			writer.Write(append([]byte(header)[:], body[:]...))

			client.Quit()
			log.Println("Mail sent successfully")
		}
	}
	return err
}