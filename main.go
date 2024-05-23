package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const url = "http://deviceshifu-plate-reader.deviceshifu.svc.cluster.local/get_measurement"

func main() {
	client := &http.Client{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "[v4] Running env test")
	})
	resp, err := client.Get(url)
	if err != nil {
		fmt.Errorf("Error: %s", err)
	} else {
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
	}

	http.ListenAndServe(":3000", nil)
}
