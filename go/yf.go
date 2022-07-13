package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
)

func urlCreater(symbol, interval, start_time, end_time string) string {
	return "https://query1.finance.yahoo.com/v8/finance/chart/" + symbol + ".T?symbol=" + symbol + ".T&period1=" + start_time + "&period2=" + end_time + "&interval=" + interval + "&includePrePost=true&events=div%7Csplit%7Cearn&lang=en-US&region=US&crumb=t5QZMhgytYZ&corsDomain=finance.yahoo.com"
}

func main() {
	// Get a greeting message and print it.
	// Allow variables to change automatically in the future
	symbol := 8304
	interval := "5m"
	start_time := "1652659200"
	end_time := "1652662800"

	url := urlCreater(strconv.Itoa(symbol), interval, start_time, end_time)

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
