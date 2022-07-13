package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func urlCreater(symbol, interval, start_time, end_time string) string {
	url :="https://query1.finance.yahoo.com/v8/finance/chart/"
	url =
	return url
}

func main() {
	// Get a greeting message and print it.
	url := "https://query1.finance.yahoo.com/v8/finance/chart/8304.T?symbol=8304.T&period1=1652659200&period2=1652662800&interval=5m&includePrePost=true&events=div%7Csplit%7Cearn&lang=en-US&region=US&crumb=t5QZMhgytYZ&corsDomain=finance.yahoo.com"
	symbol := 8304

	req, _ := http.NewRequest(http.MethodGet, url, nil)
	client := new(http.Client)
	resp, _ := client.Do(req)

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	// JSONを構造体にエンコード
	var response interface{}
	json.Unmarshal(body, &response)

	file, _ := os.Create("response.json")
	defer file.Close()

	json.NewEncoder(file).Encode(response)
}
