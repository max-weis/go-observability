package hello

import (
	"errors"
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
}

func NewService(logger zap.Logger) *service {
	return &service{logger: logger}
}

func (s *service) SayHello() *Message {
	s.logger.Info("say hello")

	return NewMessage("Hello!")
}

func (s *service) SayMessage(message string) (*Message, error) {
	if message == "" {
		s.logger.Warn("message empty")
		return nil, ErrEmptyMessage
	}
	s.logger.Info("say message", zap.String("message", message))

	return NewMessage(message), nil
}

type Message struct {
	Message string `json:"message"`
}

func NewMessage(message string) *Message {
	return &Message{Message: message}
}
