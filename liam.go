package liam

import (
	"context"
	"crypto/tls"
)

type Liam struct {
	context context.Context
	tlsConfig *tls.Config
}

// secure: false
func New() *Liam {
	return NewClient(nil)
}

func NewConfig(secure bool) *Liam {
	tlsConfig := tls.Config {
		InsecureSkipVerify: !secure,
	}

	return NewClient(&tlsConfig)
}

func NewClient(tlsConfig *tls.Config) *Liam {
	return &Liam{tlsConfig: tlsConfig}
}
// secure: false
func Smtp(server string, port int) *LSmtp {
	return New().Smtp(server, port)
}
// secure: false
func (l *Liam) Smtp(server string, port int) *LSmtp {
	return &LSmtp{
		liamog: *l,
		server: server,
		port:   port,
	}
}
// secure: false
func Send(server string, port int, username, password, emailTo, title, message string) error {
	return Smtp(server, port).
		Auth(username, password).
		AddTo(emailTo).
		Subject(title).Text(message).Send()
}