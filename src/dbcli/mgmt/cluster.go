package mgmt

import (
	"encoding/json"
	myLogger "dbcli/logger"
	"dbcli/utils"
)

type Scheduledbackup struct {
	Interval  string
	StartTime string
	Flag      string
}

//-----------------------------------

type Cluster struct {
	request utils.Request
}

func NewCluster() *Cluster {
	request := utils.NewRequest()
	return &Cluster{*request}
}

func (cluster *Cluster) ShowCluster() []byte {
	return cluster.request.SendRequest("GET", "cluster/show", nil)
}

func (cluster *Cluster) StopCluster() []byte {
	return cluster.request.SendRequest("POST", "cluster/stop", nil)
}

func (cluster *Cluster) StartCluster() []byte {
	return cluster.request.SendRequest("POST", "cluster/start", nil)
}

func (cluster *Cluster) Start(interval, starttime string) []byte {
	var scheduledbackup Scheduledbackup
	scheduledbackup.Interval = interval
	scheduledbackup.StartTime = starttime
	jsonstr, _ := json.Marshal(scheduledbackup)
	myLogger.Log.Debugf("jsonstr=", string(jsonstr))
	return cluster.request.SendRequest("POST", "cluster/scheduledbackup/start", jsonstr)
}

func (cluster *Cluster) Stop() []byte {
	return cluster.request.SendRequest("POST", "cluster/scheduledbackup/stop", nil)
}

func (cluster *Cluster) GetAll() []byte {
	return cluster.request.SendRequest("GET", "cluster/backup/all", nil)
}

func (cluster *Cluster) StartBackup() []byte {
	return cluster.request.SendRequest("POST", "cluster/backup/start", nil)
}

func (cluster *Cluster) Status() []byte {
	return cluster.request.SendRequest("GET", "cluster/backup/status", nil)
}

func (cluster *Cluster) Delete(backup_id string) []byte {
	return cluster.request.SendRequest("DELETE", "cluster/backup/"+backup_id, nil)
}

func (cluster *Cluster) DeleteAll() []byte {
	return cluster.request.SendRequest("DELETE", "cluster/backup/all", nil)
}

func (cluster *Cluster) StartRestore(backup_id string) []byte {
	return cluster.request.SendRequest("POST", "cluster/restore/"+backup_id, nil)
}



