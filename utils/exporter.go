package utils

import (
	log "github.com/Sirupsen/logrus"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	LabelALL = []string{"id", "name"}
)

func (s *SysInfo) GetExporter() error {
	ic := map[string]prometheus.CounterVec{
		"last_seen": prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: s.Type(),
				Name:      "last_seen",
				Help:      "Last time a container was seen by the exporter",
			},
			[]string{"name", "id"},
		),
		"cpu_throttled_periods_total": prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: s.Type(),
				Name:      "cpu_throttled_periods_total",
				Help:      "Number of periods with throttling.",
			},
			[]string{"name", "id"},
		),
	}
	ig := make(map[string]prometheus.GaugeVec)

	s.errors = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: s.Type(),
			Name:      "errors_total",
			Help:      "Errors while exporting container metrics.",
		},
		[]string{"component"},
	)
	s.InfoCounter = ic
	s.InfoGauge = ig
	return nil
}

func (s *SysInfo) Describe(ch chan<- *prometheus.Desc) {
	s.errors.Describe(ch)
	for _, ct := range s.InfoCounter {
		ct.Describe(ch)
	}

	for _, ig := range s.InfoGauge {
		ig.Describe(ch)
	}
}

func (s *SysInfo) Collect(ch chan<- prometheus.Metric) {
	e.Lock()
	defer e.Unlock()
	for _, ct := range s.InfoCounter {
		ct.Reset(ch)
	}
	for _, ig := range s.InfoGauge {
		ig.Reset(ch)
	}

	if err := e.collect(ch); err != nil {
		log.Errorf("Error reading container stats: %s", err)
		s.errors.WithLabelValues("collect").Inc()
	}

	for _, ct := range s.InfoCounter {
		ct.Collect(ch)
	}
	for _, ig := range s.InfoGauge {
		ig.Collect(ch)
	}
}

func (s *SysInfo) collect(ch chan<- prometheus.Metric) {

}
