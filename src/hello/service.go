package hello

import (
	"errors"
)

// ErrEmptyMessage is returned when the message is empty.
var ErrEmptyMessage = errors.New("empty message")

// Service is the interface that provides domain methods.
type Service interface {
	SayHello() *Message
	SayMessage(message string) (*Message, error)
}

// empty implementation of Service interface
type service struct {
}

func NewService() *service {
	return &service{}
}

// domain logic
func (s *service) SayHello() *Message {
	message := NewMessage("Hello!")

	return message
}

// domain logic
func (s *service) SayMessage(msg string) (*Message, error) {
	if msg == "" {
		return nil, ErrEmptyMessage
	}

	message := NewMessage(msg)
	return message, nil
}

type Message struct {
	Message string `json:"message"`
}

func NewMessage(message string) *Message {
	return &Message{Message: message}
}
