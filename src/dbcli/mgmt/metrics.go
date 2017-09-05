package mgmt

import "dbcli/utils"

type Metrics struct {
	request utils.Request
}

func NewMetrics() *Metrics {
	request := utils.NewRequest()
	return &Metrics{*request}
}

func (metrics *Metrics) Get() []byte {
	return metrics.request.SendRequest("GET", "metrics/cluster", nil)
}