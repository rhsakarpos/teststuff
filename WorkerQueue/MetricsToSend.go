package WorkerQueue

import "time"
type MetricValue struct {
	Timestamp time.Time
	MetricVal float64
}

type MetricRequest struct {
	InstanceId  string
	MetricValues []MetricValue
}

