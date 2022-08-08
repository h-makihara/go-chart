package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type return_data struct {
	Timestamp timestamp `json:"timestamp"`
	Open      open      `json:"open"`
	Close     close     `json:"close"`
	High      high      `json:"high"`
	Low       low       `json:"low"`
	Volume    volume    `json:"volume"`
}

type timestamp struct {
	name      string  `json:"name"`
	Timestamp []int64 `json:"timestamp"`
}

type open struct {
	name string    `json:"name"`
	Open []float32 `json:"open"`
}

type close struct {
	name  string    `json:"name"`
	close []float32 `json:"close"`
}

type high struct {
	name string    `json:"name"`
	High []float32 `json:"high"`
}

type low struct {
	name string    `json:"name"`
	Low  []float32 `json:"low"`
}

type volume struct {
	name   string `json:"name"`
	Volume []int  `json:"volume"`
}

type Response struct {
	Chart Chart `json:"chart"`
}
type Chart struct {
	Error  interface{} `json:"-"`
	Result []Result    `json:"result"`
}
type Result struct {
	Indicators Indicator   `json:"indicators"`
	Meta       interface{} `json:"-"`
	Timestamp  []int64     `json:"timestamp"`
}
type Indicator struct {
	Quotes []Quote `json:"quote"`
}

// Quote data is
// Open：始値
// High：最高値
// Low：最安値
// Close：終値
// Volume：取引量
type Quote struct {
	Open   []float32 `json:"open"`
	End    []float32 `json:"close"`
	High   []float32 `json:"high"`
	Low    []float32 `json:"low"`
	Volume []int     `json:"volume"`
}

// .T は日本国内株式の場合のみ追加するので、今後修正する
func urlCreater(symbol, interval, start_time, end_time string) string {
	return "https://query1.finance.yahoo.com/v8/finance/chart/" +
		symbol +
		"?symbol=" +
		symbol +
		"&period1=" +
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

func create_data(res Response, s_time int64, e_time int64) return_data {
	var resp return_data
	// name set
	resp.Timestamp.name = "timestamp"
	resp.Open.name = "open"
	resp.Close.name = "close"
	resp.High.name = "high"
	resp.Low.name = "low"
	resp.Volume.name = "volume"

	resp.Timestamp.Timestamp = res.Chart.Result[0].Timestamp

	for i := 0; i < len(res.Chart.Result); i++ {
		for j := 0; j < len(res.Chart.Result[i].Indicators.Quotes); j++ {
			for k := 0; k < len(res.Chart.Result[i].Indicators.Quotes[j].Open); k++ {
				resp.Open.Open = append(resp.Open.Open, res.Chart.Result[i].Indicators.Quotes[j].Open[k])
				resp.Close.close = append(resp.Close.close, res.Chart.Result[i].Indicators.Quotes[j].End[k])
				resp.High.High = append(resp.High.High, res.Chart.Result[i].Indicators.Quotes[j].High[k])
				resp.Low.Low = append(resp.Low.Low, res.Chart.Result[i].Indicators.Quotes[j].Low[k])
				resp.Volume.Volume = append(resp.Volume.Volume, res.Chart.Result[i].Indicators.Quotes[j].Volume[k])
			}
		}
	}
	return resp
}

func fin_print(res Response, s_time int64, e_time int64) {
	fmt.Printf("show data at %v to %v\n", time.Unix(s_time, 0), time.Unix(e_time, 0))
	for i := 0; i < len(res.Chart.Result); i++ {
		for j := 0; j < len(res.Chart.Result[i].Indicators.Quotes); j++ {
			for k := 0; k < len(res.Chart.Result[i].Indicators.Quotes[j].Open); k++ {
				fmt.Printf("%v : \n", time.Unix(res.Chart.Result[i].Timestamp[k], 0))
				fmt.Printf("  Open : %v,\t", res.Chart.Result[i].Indicators.Quotes[j].Open[k])
				fmt.Printf("End : %v,\t", res.Chart.Result[i].Indicators.Quotes[j].End[k])
				fmt.Printf("High : %v,\t", res.Chart.Result[i].Indicators.Quotes[j].High[k])
				fmt.Printf("Low : %v\t", res.Chart.Result[i].Indicators.Quotes[j].Low[k])
				fmt.Printf("Volume : %v\n\n", res.Chart.Result[i].Indicators.Quotes[j].Volume[k])
			}
		}
	}
}

func main() {
	// Get a greeting message and print it.
	// Allow variables to change automatically in the future
	//symbol := "8304"
	symbol := "SPYD"

	// interval: [15m, 30m, 60m, 90m, 1h, 1d, 5d, 1wk, 1mo, 3mo]
	interval := "30m"

	year := int(time.Now().Year())
	month := int(time.Now().Month())
	day := 2
	hour := 9
	minute := 0
	second := 0
	msecond := 0

	// set start time and end time
	s_time := unixtimeCreater(year, month, day, hour, minute, second, msecond)
	e_time := unixtimeCreater(year, month, day+1, hour+7, minute, second, msecond)

	_, t_err := strconv.Atoi(symbol)
	// Japanese ticker symbol
	if t_err == nil {
		symbol = symbol + ".T"
	}
	fmt.Printf("symbol : %v\n", symbol)

	url := urlCreater(symbol, interval, strconv.Itoa(int(s_time)), strconv.Itoa(int(e_time)))

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

	var res Response
	json.Unmarshal(body, &res)

	fin_print(res, s_time, e_time)
	var resp return_data = create_data(res, s_time, e_time)

	fmt.Println(resp)

	fmt.Printf("start time : %v\n", s_time)
	fmt.Printf("end   time : %v\n", e_time)

}
