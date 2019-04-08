package WorkerQueue

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"strconv"
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

				// do the actual sending work here
				// make a writer that produces to topic-A, using the least-bytes distribution
				w := kafka.NewWriter(kafka.WriterConfig{
					Brokers: []string{"localhost:9092"},
					Topic:   "test",
					Balancer: &kafka.LeastBytes{},
				})

				// get the string ready to be written
				var finalString = ""
				for _,metric := range work.MetricValues{
					finalString += work.InstanceId + " " + strconv.FormatFloat(metric.MetricVal,'f', 2,64) + "\n"

					w.WriteMessages(context.Background(),
						kafka.Message{
							Key:   []byte("Key-A"),
							Value: []byte(finalString),
						})
				}

				/*w.WriteMessages(context.Background(),
					kafka.Message{
						Key:   []byte("Key-A"),
						Value: []byte("Hello World!"),
					},
					kafka.Message{
						Key:   []byte("Key-B"),
						Value: []byte("One!"),
					},
					kafka.Message{
						Key:   []byte("Key-C"),
						Value: []byte("Two!"),
					},
				)*/

				w.Close()

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


