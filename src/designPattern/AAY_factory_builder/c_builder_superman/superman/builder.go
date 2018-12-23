package superman

type SuperManBuilder struct {
	superman *SuperMan
}

func NewSuperManBuilder() *SuperManBuilder {
	return &SuperManBuilder{
		new(SuperMan),
	}
}

func (s *SuperManBuilder) AndBody(body string) *SuperManBuilder {
	s.superman.Body = body
	return s
}

func (s *SuperManBuilder) AndSpecialTalent(st string) *SuperManBuilder {
	s.superman.SpecialTalent = st
	return s
}

func (s *SuperManBuilder) AndSpecialSmbol(ss string) *SuperManBuilder {
	s.superman.SpecialSmbol = ss
	return s
}

func (s *SuperManBuilder) Build() *SuperMan {
	return s.superman
}