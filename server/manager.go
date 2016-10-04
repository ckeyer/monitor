package server

import (
	"sync"

	"github.com/ckeyer/monitor/utils"
	"github.com/prometheus/client_golang/prometheus"
)

type MonitorManager struct {
	sync.RWMutex

	Monitor map[string]utils.Collector
}

func (m *MonitorManager) Init() error {
	prometheus.MustRegister(m)

	handler := prometheus.Handler()
}
