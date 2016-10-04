package server

import (
	"github.com/ckeyer/monitor/utils"
	"net/http"
	"time"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/cors"
)

const (
	API_PREFIX     = "/api"
	METRICS_PREFIX = "/metrics"
)

var manager *MonitorManager

func Serve(addr string) {
	m := martini.Classic()

	m.Use(cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Limt", "Offset", "Content-Type", "Origin", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Record-Count", "Limt", "Offset", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           time.Second * 864000,
	}))

	manager = &MonitorManager{}

	m.Group(API_PREFIX, func(r martini.Router) {
		r.Post("/metrics/:name", manager.AddNewMetrics)
	})

	m.Group(METRICS_PREFIX, func(r martini.Router) {
		r.Get("/**", manager.Metrics)
	})

	m.RunOnAddr(addr)
}

func Metrics(rw http.ResponseWriter, req *http.Request) {

}
