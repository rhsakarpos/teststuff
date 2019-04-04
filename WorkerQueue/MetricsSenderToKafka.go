package WorkerQueue

import (
	"fmt"
	"time"
)

type KafkaMetricsSender struct {
	Queue chan MetricRequest
	QuitChan chan bool
}

func (p *KafkaMetricsSender) GetMetricsSender() MetricsSenderIntf{
	sender := KafkaMetricsSender{}
	sender.Queue = make(chan MetricRequest)
	sender.QuitChan = make(chan bool)
	return &sender
}

func (p *KafkaMetricsSender) Start() {
	go func() {
		for {
			select {
			case work := <-p.Queue:
				// Receive a work request.
				fmt.Printf("GetMetricsSenderToKafka received metrics for instance %s\n and metrics %s\n", work.InstanceId, work.MetricValues)

				time.Sleep(1000)
				// do the actual sending work here

				fmt.Println("GetMetricsSenderToKafka processed metrics")

			case <-p.QuitChan:
				// We have been asked to stop.
				fmt.Println("GetMetricsSenderToKafka stopping\n")
				return
			}

		}
	}()
}
func (p *KafkaMetricsSender) Stop() {
	go func() {
		p.QuitChan <- true
	}()
}

func (p *KafkaMetricsSender) AssignMetricsToSend(request MetricRequest){
	p.Queue <- request
}


