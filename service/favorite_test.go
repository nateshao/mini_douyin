package service_test

import (
	"mini_douyin/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFavoriteAction(t *testing.T) {
	err := service.FavoriteAction(1, 1, "1")
	assert.Nil(t, err)
}

func TestFavoriteList(t *testing.T) {
	videos, err := service.FavoriteList(1)
	assert.Nil(t, err)
	assert.NotNil(t, videos)
}
