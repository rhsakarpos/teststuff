package WorkerQueue

import (
	"fmt"
	"time"
)

type PrometheusMetricsSender struct {
	Queue chan MetricRequest
	QuitChan chan bool
}

func (p *PrometheusMetricsSender) GetMetricsSender() MetricsSenderIntf{
	sender := PrometheusMetricsSender{}
	sender.Queue = make(chan MetricRequest)
	sender.QuitChan = make(chan bool)
	return &sender
}

func (p *PrometheusMetricsSender) Start() {
	go func() {
		for {
			select {
			case work := <-p.Queue:
				// Receive a work request.
				fmt.Printf("GetMetricsSenderToPrometheus received metrics for instance %s\n and metrics %s\n", work.InstanceId, work.MetricValues)

				time.Sleep(1000)
				// do the actual sending work here

				fmt.Println("GetMetricsSenderToPrometheus processed metrics")

			case <-p.QuitChan:
				// We have been asked to stop.
				fmt.Println("GetMetricsSenderToPrometheus stopping\n")
				return
			}

		}
	}()
}
func (p *PrometheusMetricsSender) Stop() {
	go func() {
		p.QuitChan <- true
	}()
}

func (p *PrometheusMetricsSender) AssignMetricsToSend(request MetricRequest){
	p.Queue <- request
}



