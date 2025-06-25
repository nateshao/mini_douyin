package service_test

import (
	"mini_douyin/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommentAction(t *testing.T) {
	comment, err := service.CommentAction(1, 1, "1", "test comment", 0)
	assert.Nil(t, err)
	assert.NotNil(t, comment)
}

func TestCommentList(t *testing.T) {
	comments, err := service.CommentList(1)
	assert.Nil(t, err)
	assert.NotNil(t, comments)
}
