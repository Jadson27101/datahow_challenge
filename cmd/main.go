package main

import (
	"datahow_challenge/config"
	"datahow_challenge/handlers"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func routerInit() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/log", handlers.AddMetric).Methods("POST")
	r.HandleFunc("/healthcheck", handlers.Healthcheck).Methods("GET")
	return r
}
func metricsRouterInit() *mux.Router {
	r := mux.NewRouter()
	r.Handle("/metrics", promhttp.Handler())
	return r
}

func main() {
	metricPort, apiPort := config.NewDefaultServerPorts()
	go func() {
		router := metricsRouterInit()
		log.Println("Start metric server on " + metricPort)
		http.ListenAndServe(":9102", router)
	}()
	router := routerInit()
	log.Println("Start main server on: " + apiPort)
	http.ListenAndServe(":5000", router)
}
