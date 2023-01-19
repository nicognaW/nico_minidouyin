package app

import (
	"context"
	"github.com/stretchr/testify/assert"
	_ "nico_minidouyin/service/feed/biz/repo"
	"testing"
)

func TestQueryFeed(t *testing.T) {
	var testUID uint32 = 0
	feed, err := QueryFeed(context.Background(), &testUID)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, testUID, feed.User.Id)
	assert.True(t, len(feed.VideoList) > 0)
	assert.True(t, len(feed.VideoList[0].Title) > 0)
}
