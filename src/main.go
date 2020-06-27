package main

import (
	"fmt"
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

	tracer, closer := config.NewTracer("localhost:5775", "observability-demo")
	defer closer.Close()

	// init domain logic
	var hs hello.Service
	hs = hello.NewService(*logger)

	// create http server
	srv := server.New(hs, *logger, tracer)

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
