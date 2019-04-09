package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"teststuff/WorkerQueue"
	"time"
)

type MetricTimeAndValue struct {
	CollectedAt time.Time `json:"collectedAt"`
	Value       string    `json:"value"`
}

type MetricsResponse struct {
	InstanceID string               `json:"instanceId"`
	MetricName string               `json:"metricName"`
	Values     []MetricTimeAndValue `json:"values"`
}

// latest+instant metrics structs begin
type InstantMetricReponseFromPrometheus struct {
	Status string `json:"status"`
	Data   Data   `json:"data"`
}
type Metric struct {
	Name     string `json:"__name__"`
	Device   string `json:"device"`
	Instance string `json:"instance"`
	Job      string `json:"job"`
}
type Result struct {
	Metric Metric        `json:"metric"`
	Value  []interface{} `json:"value"`
}
type Data struct {
	ResultType string   `json:"resultType"`
	Result     []Result `json:"result"`
}

// latest+instant metrics structs end

// range metrics struct begin
type RangeMetricReponseFromPrometheus struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric struct {
				Name     string `json:"__name__"`
				Device   string `json:"device"`
				Instance string `json:"instance"`
				Job      string `json:"job"`
			} `json:"metric"`
			Values [][]interface{} `json:"values"`
		} `json:"result"`
	} `json:"data"`
}

// range metrics struct end

func main() {

	// helloWorld()
	//metricSender()
	getLatestMetricsFromPrometheus()
	getAtInstantMetricsFromPrometheus()
	getRangeMetricsFromPrometheus()

}

func getLatestMetricsFromPrometheus() {

	response, err := http.Get("http://localhost:9090/api/v1/query?query=node_disk_bytes_read")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))

		// unmarshal the JSON response into a struct (generated using the JSON, using this https://mholt.github.io/json-to-go/
		var fv InstantMetricReponseFromPrometheus
		err0 := json.Unmarshal(data, &fv)
		fmt.Println(err0)

		metrics := make([]MetricsResponse, len(fv.Data.Result))

		// now convert to our repsonse struct, so we can marshal it and send out the JSON
		for i, res := range fv.Data.Result {
			metrics[i].InstanceID = res.Metric.Instance + res.Metric.Device
			metrics[i].MetricName = res.Metric.Name
			metrics[i].Values = make([]MetricTimeAndValue, 1)
			for _, v := range res.Value {
				switch v.(type) {
				case string:
					metrics[i].Values[0].Value = v.(string)
				case float64:
					secs := int64(v.(float64))
					nsecs := int64((v.(float64) - float64(secs)) * 1e9)

					metrics[i].Values[0].CollectedAt = time.Unix(secs, nsecs)
					//timeutil.TimestampFromFloat64(v.(float64)).Time
				/*case []interface{}:
					for i, u := range vv {
						fmt.Println(i, u)
					}*/
				default:
					fmt.Println(v, "is of a type I don't know how to handle")
				}
				/*metrics[i].CollectedAt = t
				metrics[i].Value = v*/
			}
		}

		fmt.Println("metrics struct is ", metrics)
		bArr, _ := json.Marshal(metrics)
		fmt.Println("metrics response json is ", string(bArr))

		var r1 map[string]interface{}
		err1 := json.Unmarshal(data, &r1)

		var r2 map[string]interface{} = r1["data"].(map[string]interface{})

		for k, v := range r2 {
			switch vv := v.(type) {
			case string:
				fmt.Println(k, "is string", vv)
			case float64:
				fmt.Println(k, "is float64", vv)
			case []interface{}:
				fmt.Println(k, "is an array:")
				for i, u := range vv {
					fmt.Println(i, u)
				}
			default:
				fmt.Println(k, "is of a type I don't know how to handle")
			}
		}

		//json.Unmarshal(r1["data"][1], &r)

		if err != nil || err1 != nil {
			fmt.Println(err)
		}

	}

}

func getAtInstantMetricsFromPrometheus() {

	response, err := http.Get("http://localhost:9090/api/v1/query?query=node_disk_bytes_read&time=1554720655.885")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))

		// unmarshal the JSON response into a struct (generated using the JSON, using this https://mholt.github.io/json-to-go/
		var fv InstantMetricReponseFromPrometheus
		err0 := json.Unmarshal(data, &fv)
		fmt.Println(err0)

		metrics := make([]MetricsResponse, len(fv.Data.Result))

		// now convert to our repsonse struct, so we can marshal it and send out the JSON
		for i, res := range fv.Data.Result {
			metrics[i].InstanceID = res.Metric.Instance + res.Metric.Device
			metrics[i].MetricName = res.Metric.Name
			metrics[i].Values = make([]MetricTimeAndValue, 1)
			for _, v := range res.Value {
				switch v.(type) {
				case string:
					metrics[i].Values[0].Value = v.(string)
				case float64:
					secs := int64(v.(float64))
					nsecs := int64((v.(float64) - float64(secs)) * 1e9)

					metrics[i].Values[0].CollectedAt = time.Unix(secs, nsecs)
					//timeutil.TimestampFromFloat64(v.(float64)).Time
				/*case []interface{}:
					for i, u := range vv {
						fmt.Println(i, u)
					}*/
				default:
					fmt.Println(v, "is of a type I don't know how to handle")
				}
				/*metrics[i].CollectedAt = t
				metrics[i].Value = v*/
			}
		}

		fmt.Println("metrics struct is ", metrics)
		bArr, _ := json.Marshal(metrics)
		fmt.Println("metrics response json is ", string(bArr))

		var r1 map[string]interface{}
		err1 := json.Unmarshal(data, &r1)

		var r2 map[string]interface{} = r1["data"].(map[string]interface{})

		for k, v := range r2 {
			switch vv := v.(type) {
			case string:
				fmt.Println(k, "is string", vv)
			case float64:
				fmt.Println(k, "is float64", vv)
			case []interface{}:
				fmt.Println(k, "is an array:")
				for i, u := range vv {
					fmt.Println(i, u)
				}
			default:
				fmt.Println(k, "is of a type I don't know how to handle")
			}
		}

		//json.Unmarshal(r1["data"][1], &r)

		if err != nil || err1 != nil {
			fmt.Println(err)
		}

	}

}

func getRangeMetricsFromPrometheus() {

	response, err := http.Get("http://localhost:9090/api/v1/query_range?query=node_disk_bytes_read&start=2019-04-08T06:55:30.781Z&end=2019-04-09T07:11:00.781Z&step=30")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))

		// unmarshal the JSON response into a struct (generated using the JSON, using this https://mholt.github.io/json-to-go/
		var fv RangeMetricReponseFromPrometheus
		err0 := json.Unmarshal(data, &fv)
		fmt.Println(err0)

		metrics := make([]MetricsResponse, len(fv.Data.Result))

		// now convert to our repsonse struct, so we can marshal it and send out the JSON
		for i, res := range fv.Data.Result {
			metrics[i].InstanceID = res.Metric.Instance + res.Metric.Device
			metrics[i].MetricName = res.Metric.Name
			metrics[i].Values = make([]MetricTimeAndValue, len(res.Values))
			for j1, v1 := range res.Values {
				for _, v := range v1 {
					switch v.(type) {
					case string:
						metrics[i].Values[j1].Value = v.(string)
					case float64:
						secs := int64(v.(float64))
						nsecs := int64((v.(float64) - float64(secs)) * 1e9)

						metrics[i].Values[j1].CollectedAt = time.Unix(secs, nsecs)
						//timeutil.TimestampFromFloat64(v.(float64)).Time
					/*case []interface{}:
							for i, u := range vv {
								fmt.Println(i, u)
							}*/
					default:
						fmt.Println(v, "is of a type I don't know how to handle")
					}
					/*metrics[i].CollectedAt = t
						metrics[i].Value = v*/
				}
			}
		}

		fmt.Println("metrics struct is ", metrics)
		bArr, _ := json.Marshal(metrics)
		fmt.Println("metrics response json is ", string(bArr))

		var r1 map[string]interface{}
		err1 := json.Unmarshal(data, &r1)

		var r2 map[string]interface{} = r1["data"].(map[string]interface{})

		for k, v := range r2 {
			switch vv := v.(type) {
			case string:
				fmt.Println(k, "is string", vv)
			case float64:
				fmt.Println(k, "is float64", vv)
			case []interface{}:
				fmt.Println(k, "is an array:")
				for i, u := range vv {
					fmt.Println(i, u)
				}
			default:
				fmt.Println(k, "is of a type I don't know how to handle")
			}
		}

		//json.Unmarshal(r1["data"][1], &r)

		if err != nil || err1 != nil {
			fmt.Println(err)
		}

	}

}

func helloWorld() {
	fmt.Println("hello world, enter something")
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error when reading ")
	}
	fmt.Println("after scanning " + str)
}

func metricSender() {
	// start the dispatcher
	WorkerQueue.StartDispatcher()

	// make some test metrics
	m1 := WorkerQueue.MetricValue{}
	m1.MetricVal = 11.34
	m1.Timestamp = time.Now()

	m2 := WorkerQueue.MetricValue{}
	m2.MetricVal = 12.34
	m2.Timestamp = time.Now()

	m3 := WorkerQueue.MetricValue{}
	m3.MetricVal = 13.34
	m3.Timestamp = time.Now()

	metrics := WorkerQueue.MetricRequest{}
	metrics.InstanceId = "lvm2"
	metrics.MetricValues = make([]WorkerQueue.MetricValue, 0)
	// add the two dummy values
	metrics.MetricValues = append(metrics.MetricValues, m1)

	// send them to all registered senders
	WorkerQueue.SendMetricToRegisteredSenders(metrics)

	fmt.Println("any key to continue")
	input := bufio.NewScanner(os.Stdin)
	//input.Scan()
	time.Sleep(3000 * time.Millisecond)

	metrics.MetricValues = nil
	metrics.MetricValues = make([]WorkerQueue.MetricValue, 0)
	// add one more dummy values
	metrics.MetricValues = append(metrics.MetricValues, m2)
	WorkerQueue.SendMetricToRegisteredSenders(metrics)

	fmt.Println("any key to continue")
	input = bufio.NewScanner(os.Stdin)
	//input.Scan()
	time.Sleep(3000 * time.Millisecond)

	metrics.MetricValues = nil
	metrics.MetricValues = make([]WorkerQueue.MetricValue, 0)
	// add one more dummy values
	metrics.MetricValues = append(metrics.MetricValues, m3)
	WorkerQueue.SendMetricToRegisteredSenders(metrics)

	fmt.Println("stopping senders...")

	fmt.Println("any key to exit")
	input.Scan()
}
