package hellomock

import (
	"testing"
	"github.com/golang/mock/gomock"
	"hellomock/mock"
	"github.com/stretchr/testify/mock"
)

func TestCompany_Meeting(t *testing.T) {
	person := NewPerson("王尼美")
	company := NewCompany(person)
	t.Log(company.Meeting("王尼玛"))
}

func TestCompany_Meeting2(t *testing.T) {
	ctl := gomock.NewController(t)
	mock_talker := mock_hellomock.NewMockTalker(ctl)
	mock_talker.EXPECT().SayHello(gomock.Eq("王尼玛")).Return("这是自定义的返回值，可以是任意类型。")

	company := NewCompany(mock_talker)
	t.Log(company.Meeting("王尼玛"))
	//t.Log(company.Meeting("张全蛋"))
}

type MockedPerson struct{
	mock.Mock
}

func (m *MockedPerson) SayHello(name string)(word string) {

	args := m.Called(name)
	return args.String(0)
}

func TestCompany_Meeting3(t *testing.T) {
	person := new(MockedPerson)
	person.On("SayHello", mock.Anything).Return("这是自定义的返回值")
	t.Log(person.SayHello("王尼美"))
}