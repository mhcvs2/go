package adv

type Mail struct {
	receiver, subject, context, appellation string
}

func NewMail(template *advTemplate) *Mail {
	return &Mail{
		subject:template.GetAdvSubject(),
		context:template.GetAdvContext(),
	}
}

func (m *Mail) Clone() *Mail {
	newMail := *m
	return &newMail
}

func (m *Mail)GetSubject() string {
	return m.subject
}

func (m *Mail)SetReceiver(receiver string) {
	m.receiver = receiver
}

func (m *Mail)GetReceiver() string {
	return m.receiver
}

func (m *Mail)SetAppellation(appellation string) {
	m.appellation = appellation
}

func (m *Mail)GetAppellation() string {
	return m.appellation
}
