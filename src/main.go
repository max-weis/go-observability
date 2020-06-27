package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go.uber.org/zap"
	"monitoring/config"
	"monitoring/hello"
	"monitoring/server"

	"net/http"
)

func main() {
	logger, _ := initLogger()
	defer logger.Sync()

	config := config.NewConfig(*logger)

	logger.Info("start server", zap.String("port", config.Port))

	tracer, closer := config.NewTracer()
	if closer != nil {
		defer closer.Close()
	}

	// init domain logic
	var hs hello.Service
	hs = hello.NewService()
	hs = hello.NewLoggingService(*logger, hs)

	if tracer != nil {
		hs = hello.NewTracingService(tracer, hs)
	}

	hs = hello.NewInstrumentingService(
		promauto.NewCounter(
			prometheus.CounterOpts{
				Name: "say_hello_total",
				Help: "The total number of said hellos",
			}),
		promauto.NewCounter(
			prometheus.CounterOpts{
				Name: "say_message_total",
				Help: "The total number of said messages",
			}),
		hs)

	// create http server
	srv := server.New(hs, tracer)

	logger.Info("server can accept requests")
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", config.Port), srv)
}

func initLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"./log/monitoring.log",
		"stdout",
	}
	return cfg.Build()
}
