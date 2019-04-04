package main

import (
	"bufio"
	"fmt"
	"os"
	"teststuff/WorkerQueue"
	"time"
)

func main(){

	/*fmt.Println("hello world, enter something")
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil{
		fmt.Println("error when reading ")
	}
	fmt.Println("after scanning " + str)*/

	// start the dispatcher
	WorkerQueue.StartDispatcher()

	// make some test metrics
	m1 := WorkerQueue.MetricValue{}
	m1.MetricVal = 12.34
	m1.Timestamp = time.Now()

	m2 := WorkerQueue.MetricValue{}
	m2.MetricVal = 13.34
	m2.Timestamp = time.Now()

	m3 := WorkerQueue.MetricValue{}
	m3.MetricVal = 14.34
	m3.Timestamp = time.Now()

	metrics := WorkerQueue.MetricRequest{}
	metrics.InstanceId = "1"
	metrics.MetricValues = make([]WorkerQueue.MetricValue, 2)
	// add the two dummy values
	metrics.MetricValues = append(metrics.MetricValues, m1, m2)


	// send them to all registered senders
	WorkerQueue.SendMetricToRegisteredSenders(metrics)

	fmt.Println("any key to continue")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	metrics.MetricValues = nil
	metrics.MetricValues = make([]WorkerQueue.MetricValue, 1)
	// add one more dummy values
	metrics.MetricValues = append(metrics.MetricValues, m3)
	WorkerQueue.SendMetricToRegisteredSenders(metrics)


	fmt.Println("stopping senders...")


	fmt.Println("any key to exit")
	input.Scan()
}
