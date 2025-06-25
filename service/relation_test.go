package service_test

import (
	"mini_douyin/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRelationAction(t *testing.T) {
	err := service.RelationAction(1, 2, "1")
	assert.Nil(t, err)
}

func TestFollowList(t *testing.T) {
	users, err := service.FollowList(1)
	assert.Nil(t, err)
	assert.NotNil(t, users)
}

func TestFollowerList(t *testing.T) {
	users, err := service.FollowerList(1)
	assert.Nil(t, err)
	assert.NotNil(t, users)
}
