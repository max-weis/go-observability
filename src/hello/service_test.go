// +build unit

package hello

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMessage(t *testing.T) {
	tests := []struct {
		have string
		want *Message
	}{
		{"test", &Message{Message: "test"}},
		{"", &Message{Message: ""}},
	}
	for _, tt := range tests {
		got := NewMessage(tt.have)
		assert.Equal(t, got, tt.want, fmt.Sprintf("SayMessage() got = %v, want %v", got, tt.want))
	}
}

func TestService_SayHello(t *testing.T) {
	s := service{}

	message := s.SayHello()

	assert.Equal(t, "Hello!", message.Message)
}

func TestService_SayMessage(t *testing.T) {
	s := service{}

	message, err := s.SayMessage("test")

	assert.Nil(t, err)
	assert.Equal(t, "test", message.Message)
}

func TestService_SayMessageError(t *testing.T) {
	s := service{}

	_, err := s.SayMessage("")

	assert.Equal(t, ErrEmptyMessage, err)
}
