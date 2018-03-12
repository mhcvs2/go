package adv

type advTemplate struct {
	advSubject string
	advContext string
}

func NewAdvTemplate() *advTemplate {
	return &advTemplate{
		advSubject:"抽奖活动",
		advContext:"抽奖活动详情",
	}
}

func (a advTemplate) GetAdvSubject() string {
	return a.advSubject
}

func (a advTemplate) GetAdvContext() string {
	return a.advContext
}
