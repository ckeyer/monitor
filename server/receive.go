package server

import (
	"net/http"
)

func (m *MonitorManager) Receive(rw http.ResponseWriter, req *http.Request) {

	m.RLock()
	defer m.RUnlock()

}

func (m *MonitorManager) AddNewMetrics(rw http.ResponseWriter, req *http.Request) {

	m.Lock()
	defer m.Unlock()

}
