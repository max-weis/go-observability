package hello

import (
	"github.com/opentracing/opentracing-go"
)

type tracingService struct {
	tracer opentracing.Tracer
	next   Service
}

func NewTracingService(tracer opentracing.Tracer, next Service) *tracingService {
	return &tracingService{
		tracer: tracer,
		next:   next,
	}
}

func (s *tracingService) SayHello() *Message {
	defer func() {
		s.tracer.StartSpan("say_hello_span")
	}()

	return s.next.SayHello()
}

func (s *tracingService) SayMessage(message string) (*Message, error) {
	defer func() {
		s.tracer.StartSpan("say_message_span")
	}()

	return s.next.SayMessage(message)
}
