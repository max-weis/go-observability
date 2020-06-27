package hello

import (
	"errors"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

// ErrEmptyMessage is returned when the message is empty.
var ErrEmptyMessage = errors.New("empty message")

// Service is the interface that provides domain methods.
type Service interface {
	SayHello() *Message
	SayMessage(message string) (*Message, error)
}

// empty implementation
type service struct {
	logger zap.Logger
	tracer opentracing.Tracer
}

func NewService(logger zap.Logger, tracer opentracing.Tracer) *service {
	return &service{logger: logger, tracer: tracer}
}

func (s *service) SayHello() *Message {
	s.logger.Info("say hello")

	span := s.tracer.StartSpan("say_hello_span")
	defer span.Finish()

	return NewMessage("Hello!")
}

func (s *service) SayMessage(message string) (*Message, error) {
	s.logger.Info("say message", zap.String("message", message))

	span := s.tracer.StartSpan("say_message_span")
	defer span.Finish()

	if message == "" {
		s.logger.Warn("message empty")
		return nil, ErrEmptyMessage
	}

	return NewMessage(message), nil
}

type Message struct {
	Message string `json:"message"`
}

func NewMessage(message string) *Message {
	return &Message{Message: message}
}
