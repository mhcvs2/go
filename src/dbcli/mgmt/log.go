package mgmt

import "dbcli/utils"

type ClusterLog struct {
	request utils.Request
}

func NewClusterLog() *ClusterLog {
	request := utils.NewRequest()
	return &ClusterLog{*request}
}

func (clusterlog *ClusterLog) ClusterLog() []byte {
	return clusterlog.request.SendRequest("GET", "log/cluster", nil)
}

func (clusterlog *ClusterLog) GetDockerContainerLog(containerId string) []byte {
	return clusterlog.request.SendRequest("GET", "log/container/"+containerId, nil)
}


