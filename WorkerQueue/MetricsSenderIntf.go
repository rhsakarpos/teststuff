package WorkerQueue

type MetricsSenderIntf interface {

	//GetMetricsSender(queue chan MetricRequest, quitChan chan bool) MetricsSenderIntf
	GetMetricsSender() MetricsSenderIntf
	AssignMetricsToSend(request MetricRequest)
	Start()
	Stop()
}

