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
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	config := config.NewConfig(*logger)

	logger.Info("start server", zap.String("port", config.Port))

	var hs hello.Service
	hs = hello.NewService(*logger)

	srv := server.New(hs, *logger)

	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", config.Port), srv)
}
