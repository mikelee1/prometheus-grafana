package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/hnw/go-smc/smc"
)

// Exporter defines the metrics types
type Exporter struct {
	counter    prometheus.Counter
	counterVec prometheus.CounterVec
	cpuTemp    prometheus.Gauge
}

// NewExporter returns a custom exporter
func NewExporter(metricsPrefix string) *Exporter {
	counter := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: metricsPrefix,
		Name:      "counter_metric",
		Help:      "This is a counter for number of total api calls"})

	counterVec := *prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: metricsPrefix,
		Name:      "counter_vec_metric",
		Help:      "This is a counter vec for number of all api calls"},
		[]string{"endpoint"})
	cpuTemp := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: metricsPrefix,
		Name:      "cpu_temperature_celsius",
		Help:      "Current temperature of the CPU.",
	})

	cpuTemp.Set(65.3)

	return &Exporter{
		counter:    counter,
		counterVec: counterVec,
		cpuTemp:    cpuTemp,
	}
}

// Collect implements prometheus.Collector.Collect
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	e.counter.Collect(ch)
	e.counterVec.Collect(ch)
	e.cpuTemp.Collect(ch)
}

// Describe implements prometheus.Collector.Describe
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	e.counter.Describe(ch)
	e.counterVec.Describe(ch)
	e.cpuTemp.Describe(ch)
}

// IncreCounter increments counter
func (e *Exporter) IncreCounter() {
	e.counter.Inc()
}

// IncreCounterWithEndpoint increments counter
func (e *Exporter) IncreCounterWithEndpoint(endpoint string) {
	e.counterVec.With(prometheus.Labels{"endpoint": endpoint}).Inc()
}

// IncreCounterWithEndpoint increments counter
func (e *Exporter) ChangeCpuTemp() {
	smc.Open()
	defer smc.Close()
	e.cpuTemp.Set(smc.GetTemperature(smc.KEY_CPU_TEMP))
}
