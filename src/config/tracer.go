package config

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"go.uber.org/zap"
	"io"
)

func (c *Config) NewTracer() (opentracing.Tracer, io.Closer) {
	cfg, err := config.FromEnv()
	if err != nil {
		c.logger.Warn("could not get jaeger env vars", zap.Error(err))
	}
	tracer, closer, err := cfg.NewTracer(
		config.Logger(jaeger.StdLogger),
	)
	if err != nil {
		c.logger.Warn("could not init tracer", zap.Error(err))
		return nil, nil
	}

	return tracer, closer
}
