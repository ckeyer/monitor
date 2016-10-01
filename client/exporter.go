package client

import (
	"math/rand"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

type Exporter struct {
	mutex sync.RWMutex

	errors           *prometheus.CounterVec
	cpuUsageSeconds  *prometheus.CounterVec
	memoryUsageBytes *prometheus.GaugeVec
	pickCount        *prometheus.CounterVec
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	e.errors.Describe(ch)

	e.cpuUsageSeconds.Describe(ch)
	e.memoryUsageBytes.Describe(ch)
	e.pickCount.Describe(ch)
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	e.mutex.Lock() // To protect metrics from concurrent collects.
	defer e.mutex.Unlock()
	e.cpuUsageSeconds.Reset()
	e.memoryUsageBytes.Reset()
	e.pickCount.Reset()

	name := "container.Name"
	id := "container.ID"
	image := "container.Image"

	tmp := e.cpuUsageSeconds.WithLabelValues([]string{name, id, image, "kernel"}...)
	tmp.Set(3.14)
	e.cpuUsageSeconds.WithLabelValues([]string{name, id, image, "user"}...).Set(6.28)

	// Memory stats
	r := rand.Float64() * 40
	e.memoryUsageBytes.WithLabelValues([]string{name, id, image}...).Set(60 + r)

	e.pickCount.WithLabelValues([]string{name, id, image}...).Inc()

	e.errors.Collect(ch)
	e.cpuUsageSeconds.Collect(ch)
	e.memoryUsageBytes.Collect(ch)
	e.pickCount.Collect(ch)
}
