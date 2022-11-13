package main

import "github.com/prometheus/client_golang/prometheus"

var (
	// TxPwr is the transmission power in dBmV from the modem to the TAP
	TxPwr = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "fibertel",
			Name:      "tx_pwr",
			Help:      "TX power in dBmV",
		})
	// TxFreq is the selected channel to communicate within the TAP
	TxFreq = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "fibertel",
			Name:      "tx_freq",
			Help:      "TX freq in MHz",
		})
	// RxPwr is the transmission power in dBmV from the TAP to the modem
	RxPwr = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "fibertel",
			Name:      "rx_pwr",
			Help:      "RX power in dBmV",
		})
	// RxFreq is the selected channel to communicate within the modem
	RxFreq = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "fibertel",
			Name:      "rx_freq",
			Help:      "RX freq in MHz",
		})
	// Mer is digital complex baseband signal-to-noise ratio (SNR)
	Mer = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "fibertel",
			Name:      "mer",
			Help:      "MER power in dB",
		})
)
