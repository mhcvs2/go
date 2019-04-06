package commands

import (
	"fmt"
	"strings"
)

const (
	DIVIDE_FLAG = " "
	PREFIX = "-"
)

type CommandVO struct {
	commandName string
	paramList []string
	dataList []string
}

func NewCommandVO(commandStr string) *CommandVO {
	c := CommandVO{
		commandName: "",
		paramList: make([]string, 0),
		dataList: make([]string, 0),
	}
	if commandStr != "" {
		complextStr := strings.Split(commandStr, DIVIDE_FLAG)
		c.commandName = complextStr[0]
		for i := 1; i < len(complextStr); i++ {
			str := complextStr[i]
			if strings.Index(str, PREFIX) == 0 {
				c.paramList = append(c.paramList, strings.TrimSpace(strings.Replace(str, PREFIX, "", -1)))
			} else {
				c.dataList = append(c.dataList, strings.TrimSpace(str))
			}
		}

	} else {
		fmt.Println("failed to parse command")
	}
	return &c
}

func (s *CommandVO)CommandName() string {
    return s.commandName
}

func (s *CommandVO)ParamList() []string {
    if len(s.paramList) == 0 {
    	s.paramList = append(s.paramList, "")
	}
    return s.paramList
}

func (s *CommandVO)FormatData() string {
    return strings.Join(s.dataList, " ")
}
