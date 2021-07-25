package main

import "github.com/prometheus/client_golang/prometheus"

var (
	TxPwr = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "fibertel",
			Name:      "tx_pwr",
			Help:      "TX power in dBmV",
		})
	TxFreq = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "fibertel",
			Name:      "tx_freq",
			Help:      "TX freq in MHz",
		})
	RxPwr = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "fibertel",
			Name:      "rx_pwr",
			Help:      "RX power in dBmV",
		})
	RxFreq = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "fibertel",
			Name:      "rx_freq",
			Help:      "RX freq in MHz",
		})
	Mer = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "fibertel",
			Name:      "mer",
			Help:      "MER power in dB",
		})
)
