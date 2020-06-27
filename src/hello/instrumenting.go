package hello

import (
	"github.com/prometheus/client_golang/prometheus"
)

type instrumentingService struct {
	helloCounter   prometheus.Counter
	messageCounter prometheus.Counter
	next           Service
}

func NewInstrumentingService(helloCounter prometheus.Counter, messageCounter prometheus.Counter, s Service) Service {
	return &instrumentingService{
		helloCounter:   helloCounter,
		messageCounter: messageCounter,
		next:           s,
	}
}

func (s *instrumentingService) SayHello() *Message {
	defer func() {
		s.helloCounter.Inc()
	}()

	return s.next.SayHello()
}

func (s *instrumentingService) SayMessage(message string) (*Message, error) {
	defer func() {
		s.messageCounter.Inc()
	}()

	return s.next.SayMessage(message)
}
