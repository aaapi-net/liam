package liam

import "fmt"

type LSmtp struct {
	liamog Liam

	server string
	port int

	username string
	password string
}

func (ls *LSmtp) Auth(username string, password string) *LMail  {
	ls.username = username
	ls.password = password
	return &LMail{
		smtp: ls,
	}
}

func (ls *LSmtp) getAddr() string  {
	return fmt.Sprintf("%s:%d",  ls.server, ls.port)
}