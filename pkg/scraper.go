package main

import (
	"fmt"
	"github.com/prometheus/common/expfmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:9090/metrics")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	var textMetricParser expfmt.TextParser
	metricFamilyMap, err := textMetricParser.TextToMetricFamilies(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	for k, v := range metricFamilyMap {
		fmt.Println("KEY: ", k)
		fmt.Println("VAL: ", v)
	}
}
