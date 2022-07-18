package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

// chart data is
// Open：その日の始値
// High：その日の最高値
// Low：その日の最安値
// Close：その日の終値
// Volume：その日の取引量
/*
type chart struct {
	c_open    uint
	c_high    uint
	c_low     uint
	c_close   uint
	c_volume  uint
	timestamp uint64
}
*/

func urlCreater(symbol, interval, start_time, end_time string) string {
	return "https://query1.finance.yahoo.com/v8/finance/chart/" +
		symbol +
		".T?symbol=" +
		symbol +
		".T&period1=" +
		start_time +
		"&period2=" +
		end_time +
		"&interval=" +
		interval +
		"&includePrePost=true&events=div%7Csplit%7Cearn&lang=en-US&region=US&crumb=t5QZMhgytYZ&corsDomain=finance.yahoo.com"
}

func unixtimeCreater(yy, mm, dd, hh, min, sec, msec int) int64 {
	m_name := []time.Month{time.January, time.February, time.March, time.April, time.May, time.June, time.July, time.August, time.September, time.October, time.November, time.December}

	ut := time.Date(yy, m_name[mm-1], dd, hh, min, sec, msec, time.UTC).Unix()
	return ut
}

func main() {
	// Get a greeting message and print it.
	// Allow variables to change automatically in the future
	symbol := 8304
	interval := "5m"
	year := int(time.Now().Year())
	month := int(time.Now().Month())
	day := 14
	hour := 9
	minute := 0
	second := 0
	msecond := 0

	// set start time and end time
	s_time := unixtimeCreater(year, month, day, hour, minute, second, msecond)
	fmt.Println(s_time)
	e_time := unixtimeCreater(year, month, day+1, hour+6, minute, second, msecond)
	fmt.Println(e_time)

	url := urlCreater(strconv.Itoa(symbol), interval, strconv.Itoa(int(s_time)), strconv.Itoa(int(e_time)))

	get_res, h_err := http.Get(url)
	if h_err != nil {
		fmt.Println(h_err)
		return
	}
	defer get_res.Body.Close()

	body, err := io.ReadAll(get_res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	// JSONを構造体にエンコード
	var response map[string]interface{}
	json.Unmarshal(body, &response)
	fmt.Println(string(body))

	// file out
	/*
		file, _ := os.Create("response2.json")
		defer file.Close()

		json.NewEncoder(file).Encode(response)
	*/
}
