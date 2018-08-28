package handler

import "fmt"

type IWoman interface {
	GetType() LevelType
	GetRequest() string
}

type Woman struct {
	ltype LevelType
	request string
}

func NewWoman(ltype LevelType, request string) *Woman {
	wm := &Woman{ltype:ltype}
	switch ltype {
	case FATHER:
		wm.request = fmt.Sprintf("daughter's request is: %s", request)
	case HUSBAND:
		wm.request = fmt.Sprintf("wife's request is: %s", request)
	case SON:
		wm.request = fmt.Sprintf("mother's request is: %s", request)
	}
	return wm
}

func (w *Woman) GetType() LevelType {
	return w.ltype
}

func (w *Woman) GetRequest() string {
	return w.request
}
