package liam

import (
	"errors"
	"fmt"
	"net/mail"
	"strings"
)

type LMail struct {
	smtp *LSmtp

	sender string
	to []string
	cc      []string		// TODO: BCC messages by including in the to parameter but not including it in the msg headers
	subject string
	body interface{}
}

func (lm *LMail) Text(body interface{}) *MailText {

	lm.body = body
	return &MailText{*lm}
}

func (lm *LMail) Struct(body interface{}) *MailStruct {

	lm.body = body
	return &MailStruct{MailText{*lm}}
}
/*
var body = map[string]interface{}{"Username": "Dave", "Type": "a new client"}

var template = `Hello {{ .Username }}, your are {{ .Type }}.`

The name of a key of the data, which must be a map, preceded
  by a period, such as
    .Key
  The result is the map element value indexed by the key.
  Key invocations may be chained and combined with fields to any
  depth:
    .Field1.Key1.Field2.Key2
  Although the key must be an alphanumeric identifier, unlike with
  field names they do not need to start with an upper case letter.
  Keys can also be evaluated on variables, including chaining:
    $x.key1.key2
 */
func (lm *LMail) Template(body map[string]interface{}, template string) *MailTemplate {
	lm.body = body
	return &MailTemplate{mail: *lm, template: template}
}

func (lm *LMail) Sender(sender string) *LMail {
	lm.sender = sender
	return lm
}

func (lm *LMail) Subject(subject string) *LMail {
	lm.subject = subject
	return lm
}

func (lm *LMail) AddTo(to ...string) *LMail {
	lm.to = append(lm.to, to...)
	return lm
}

func (lm *LMail) SetTo(to []string) *LMail {
	lm.to = to
	return lm
}


func (lm *LMail) AddCopyTo(cc ...string) *LMail {
	lm.cc = append(lm.cc, cc...)
	return lm
}

func (lm *LMail) SetCopyTo(cc []string) *LMail {
	lm.cc = cc
	return lm
}

func (lm *LMail) bodyHeader() (message string, err error) {
	if lm.sender == "" {
		lm.sender = lm.smtp.username
	}

	message += fmt.Sprintf("From: %s\r\n", lm.sender)
	if len(lm.to) > 0 {
		if !isValidEmails(lm.to) {
			return "", errors.New("wrong email address to")
		}
		message += fmt.Sprintf("To: %s\r\n", strings.Join(lm.to, ";"))
	} else {
		return "", errors.New("set mail to")
	}

	if len(lm.cc) > 0 {
		if !isValidEmails(lm.cc) {
			return "", errors.New("wrong email address cc")
		}
		message += fmt.Sprintf("Cc: %s\r\n", strings.Join(lm.cc, ";"))
	}

	message += fmt.Sprintf("Subject: %s\r\n", lm.subject)
	message += "\r\n\r\n" // + body

	return
}


func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func isValidEmails(emails []string) bool  {
	for _, email := range emails {
		if !isValidEmail(email) {
			return false
		}
	}
	return true
}