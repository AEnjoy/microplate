package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const url = "http://deviceshifu-plate-reader.deviceshifu.svc.cluster.local/get_measurement"

func main() {
	client := &http.Client{}
	ticker := time.NewTicker(3 * time.Second)
	for range ticker.C {
		resp, err := client.Get(url)
		if err != nil {
			fmt.Errorf("error: %s\n", err)
		} else {
			body, _ := ioutil.ReadAll(resp.Body)
			str := strings.Split(string(body), "\n")
			count := 0
			value := 0.0
			for _, s := range str {
				numStr := strings.Split(s, " ")
				for _, s2 := range numStr {
					float, err := strconv.ParseFloat(s2, 64)
					if err != nil {
						continue
					}
					count++
					value += float
				}
			}
			if count != 0 {
				fmt.Printf("%.2f\n", value/float64(count))
			}
			resp.Body.Close()
		}
	}
}
