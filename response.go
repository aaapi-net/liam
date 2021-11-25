package liam

import (
	"crypto/tls"
	"errors"
	"log"
	"net/smtp"
	"strings"
)

type response interface {
	Send() error
	getMail() *LMail
	getBodyBytes() ([]byte, error)
}

func getResponse(r response)  error {
	header, err := r.getMail().bodyHeader()
	mail := r.getMail()

	if err == nil {
		body, err := r.getBodyBytes()
		auth := smtp.PlainAuth("", mail.smtp.username, mail.smtp.password, mail.smtp.server)

		log.Println("Header: ", header, string(body))

		if err == nil {
			if mail.smtp.liamog.tlsConfig == nil {
				err = smtp.SendMail(mail.smtp.getAddr(), auth, mail.sender, mail.to, append([]byte(header)[:], body[:]...))
			} else {
				err = sendMailTLS(mail.smtp, auth, mail.sender, mail.to, append([]byte(header)[:], body[:]...))
			}
			if err != nil {
				return err
			}
			log.Println("Mail sent successfully")
		}
	}
	return err
}


func sendMailTLS(lsmtp *LSmtp, auth smtp.Auth, from string, to []string, msg []byte) error {

	if err := validateLine(from); err != nil {
		return err
	}
	for _, recp := range to {
		if err := validateLine(recp); err != nil {
			return err
		}
	}
	conn, err := tls.Dial("tcp", lsmtp.getAddr(), lsmtp.liamog.tlsConfig)
	if err != nil {
		return err
	}
	defer conn.Close()
	c, err := smtp.NewClient(conn, lsmtp.server)
	if err != nil {
		return err
	}
	defer c.Close()
	if err = c.Hello("localhost"); err != nil {
		return err
	}
	if err = c.Auth(auth); err != nil {
		return err
	}
	if err = c.Mail(from); err != nil {
		return err
	}
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(msg)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}

// validateLine checks to see if a line has CR or LF as per RFC 5321
func validateLine(line string) error {
	if strings.ContainsAny(line, "\n\r") {
		return errors.New("a line must not contain CR or LF")
	}
	return nil
}
