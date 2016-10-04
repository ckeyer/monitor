package client

import (
	"net/http"

	"github.com/ckeyer/monitor/utils"
)

type SysInfo utils.SysInfo

func Handle(rw http.ResponseWriter, req *http.Request) {

}

// func Cli() {
// 	exporter := &Exporter{
// 		errors: prometheus.NewCounterVec(prometheus.CounterOpts{
// 			Namespace: namespace,
// 			Name:      "errors_total",
// 			Help:      "Errors while exporting container metrics.",
// 		},
// 			[]string{"component"},
// 		),
// 		cpuUsageSeconds: prometheus.NewCounterVec(prometheus.CounterOpts{
// 			Namespace: namespace,
// 			Name:      "cpu_usage_seconds_total",
// 			Help:      "Total seconds of cpu time consumed.",
// 		},
// 			[]string{"name", "id", "type"},
// 		),

// 		memoryUsageBytes: prometheus.NewGaugeVec(prometheus.GaugeOpts{
// 			Namespace: namespace,
// 			Name:      "memory_max_usage_bytes",
// 			Help:      "Maximum memory usage ever recorded in bytes.",
// 		},
// 			[]string{"name", "id"},
// 		),
// 	}
// 	prometheus.MustRegister(exporter)

// 	log.Printf("Starting Server: %s", ":8000")
// 	handler := prometheus.Handler()
// 	http.Handle("/state", handler)

// 	log.Fatal(http.ListenAndServe(":8000", nil))
// }
