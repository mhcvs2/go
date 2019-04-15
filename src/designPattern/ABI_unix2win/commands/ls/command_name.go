package ls

import "designPattern/ABI_unix2win/commands"

const (
	DEFAULT_PARAM = ""
	A_PARAM = "a"
	L_PARAM = "l"
)

type LS struct {
	*commands.CommandName
}

func NewLS() *LS {
	res := &LS{}
	res.CommandName = commands.NewCommandName(res)
	return res
}

func (c *LS) OperateParam() string {
	return DEFAULT_PARAM
}

func (c *LS) Echo(vo commands.CommandVO) string {
	return ls(vo.FormatData())
}

//-----------------------------------------
type LSA struct {
	*commands.CommandName
}

func NewLSA() *LSA {
	res := &LSA{}
	res.CommandName = commands.NewCommandName(res)
	return res
}

func (c *LSA) OperateParam() string {
	return A_PARAM
}

func (c *LSA) Echo(vo commands.CommandVO) string {
	return ls_a(vo.FormatData())
}

//-------------------------------------------
type LSL struct {
	*commands.CommandName
}

func NewLSL() *LSL {
	res := &LSL{}
	res.CommandName = commands.NewCommandName(res)
	return res
}

func (c *LSL) OperateParam() string {
	return L_PARAM
}

func (c *LSL) Echo(vo commands.CommandVO) string {
	return ls_l(vo.FormatData())
}
