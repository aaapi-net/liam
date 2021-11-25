package liam

type MailStruct struct {
	MailText
}

func (ms *MailStruct) AsJson() *MailJson {
	return &MailJson{ms.mail}
}

func (ms *MailStruct) AsTemplate(template string) *MailTemplate {
	return &MailTemplate{ms.mail, template}
}