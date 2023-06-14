package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Signal is the structure coming from the modem stats
type Signal struct {
	TxPwr  float64
	TxFreq float64
	RxPwr  float64
	RxFreq float64
	Mer    float64
}

func main() {
	signal := Signal{}

	go func() {
		for {

			s, err := signal.getCurrentValues()

			if err != nil {
				log.Println(err.Error())
			}
			TxPwr.Set(s.TxPwr)
			TxFreq.Set(s.TxFreq)
			RxPwr.Set(s.RxPwr)
			RxFreq.Set(s.RxFreq)
			Mer.Set(s.Mer)

			time.Sleep(10 * time.Second)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	prometheus.MustRegister(TxPwr)
	prometheus.MustRegister(TxFreq)
	prometheus.MustRegister(RxPwr)
	prometheus.MustRegister(RxFreq)
	prometheus.MustRegister(Mer)
	err := http.ListenAndServe(":9100", nil)
	if err != nil {
		log.Println(err.Error())
	}

}

func (s *Signal) getCurrentValues() (Signal, error) {
	var signal Signal

	url := "http://dameunaip.com.ar/asp/nivelesprima.asp"

	resp, err := makeRequestWithExponentialBackoff(url)

	if err != nil {
		return signal, err
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return signal, err
	}

	bodyString := string(bodyBytes)

	lines := strings.Split(bodyString, "\n")

	values := make([]string, 0)

	for _, line := range lines {
		if strings.Contains(line, "valor") {
			line = strings.Split(line, ">")[1]
			line = strings.Split(line, " ")[0]
			line = strings.Replace(line, ",", ".", 1)
			_, err := strconv.ParseFloat(line, 64)

			if err == nil {
				values = append(values, line)
			}
		}
	}

	if len(values) != 5 {
		return signal, err
	}

	signal.TxPwr, _ = strconv.ParseFloat(values[0], 32)
	signal.TxFreq, _ = strconv.ParseFloat(values[1], 32)
	signal.RxPwr, _ = strconv.ParseFloat(values[2], 32)
	signal.RxFreq, _ = strconv.ParseFloat(values[3], 32)
	signal.Mer, _ = strconv.ParseFloat(values[4], 32)
	return signal, nil

}

func makeRequestWithExponentialBackoff(url string) (*http.Response, error) {
	maxRetries := 5
	retryDelay := 1 * time.Second

	client := http.Client{
		Timeout: 15 * time.Second,
	}

	for i := 0; i < maxRetries; i++ {
		resp, err := client.Get(url)
		if err == nil {
			return resp, nil
		}

		delay := retryDelay * time.Duration(1<<uint(i))
		time.Sleep(delay)
	}

	return nil, nil
}
