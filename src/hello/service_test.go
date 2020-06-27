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
