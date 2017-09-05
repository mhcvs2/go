package mgmt

import (
	"encoding/json"
	myLogger "dbcli/logger"
	"dbcli/utils"
)

type Database struct {
	Name string
}

//------------------------

type Db struct {
	request utils.Request
}

func NewDb() *Db {
	request := utils.NewRequest()
	return &Db{*request}
}

func (db *Db) Post(name string) []byte {
	var database Database
	database.Name = name
	jsonstr, _ := json.Marshal(database)
	myLogger.Log.Debugf("jsonstr=", string(jsonstr))
	return db.request.SendRequest("POST", "database", jsonstr)
}

func (db *Db) Delete(name string) []byte {
	return db.request.SendRequest("POST", "database/"+name, nil)
}