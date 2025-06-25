package service_test

import (
	"mini_douyin/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMessageAction(t *testing.T) {
	err := service.MessageAction(1, 2, "1", "hi")
	assert.Nil(t, err)
}

func TestMessageChat(t *testing.T) {
	messages, err := service.MessageChat(1, 2)
	assert.Nil(t, err)
	assert.NotNil(t, messages)
}
