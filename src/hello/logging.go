package hello

import "go.uber.org/zap"

type loggingService struct {
	logger zap.Logger
	next   Service
}

func NewLoggingService(logger zap.Logger, next Service) *loggingService {
	return &loggingService{
		logger: logger,
		next:   next,
	}
}

func (s *loggingService) SayHello() *Message {
	defer func() {
		s.logger.Info("say hello")
	}()

	return s.next.SayHello()
}

func (s *loggingService) SayMessage(message string) (*Message, error) {
	defer func() {
		s.logger.Info("say message", zap.String("message", message))
	}()

	return s.next.SayMessage(message)
}
