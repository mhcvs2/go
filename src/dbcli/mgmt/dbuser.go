package mgmt

import (
	"encoding/json"
	myLogger "dbcli/logger"
	"dbcli/utils"
)

type Dbuser struct {
	Dbname   string
	Username string
	Pwd      string
}

//------------------------

type Dbu struct {
	request utils.Request
}

func NewDbu() *Dbu {
	request := utils.NewRequest()
	return &Dbu{*request}
}

func (dbu *Dbu) Post2(dbName, userName, password string) []byte {
	var dbuser Dbuser
	dbuser.Dbname = dbName
	dbuser.Username = userName
	dbuser.Pwd = password
	jsonstr, _ := json.Marshal(dbuser)
	myLogger.Log.Debugf("jsonstr=", string(jsonstr))
	return dbu.request.SendRequest("POST", "dbuser", jsonstr)
}

func (dbu *Dbu) Delete2(dbName, userName string) []byte {
	return dbu.request.SendRequest("POST", "dbuser/"+dbName+"/"+userName, nil)
}