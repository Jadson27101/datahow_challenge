package handlers

import (
	"context"
	"datahow_challenge/config"
	"datahow_challenge/models"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"net/http"
)

var ctx = context.Background()
var (
	OpsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "unique_ip_addresses",
		Help: "The total number of unique user ip",
	})
)

func AddMetric(w http.ResponseWriter, r *http.Request) {
	info := &models.Info{}
	if err := json.NewDecoder(r.Body).Decode(info); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid body"))
	}
	if err := info.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid body"))
	}
	redisClient := config.NewDefaultRedisClient()
	_, err := redisClient.Get(ctx, info.Ip).Result()
	if err == redis.Nil {
		err := redisClient.Set(ctx, info.Ip, true, 0).Err()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		OpsProcessed.Inc()
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Redis unavailable"))
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
