package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const url = "http://deviceshifu-plate-reader.deviceshifu.svc.cluster.local/get_measurement"

func main() {
	client := &http.Client{}

	resp, _ := client.Get(url)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	str := strings.Split(string(body), "\n")
	count := 0
	value := 0.0
	for _, s := range str {
		numStr := strings.Split(s, " ")
		for _, s2 := range numStr {
			float, err := strconv.ParseFloat(s2, 64)
			if err != nil {
				return
			}
			count++
			value += float
		}
	}
	if count != 0 {
		fmt.Printf("%.2f", value/float64(count))
	}
	select {}
}
