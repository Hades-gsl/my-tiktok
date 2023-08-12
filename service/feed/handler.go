package main

import (
	"context"
	feed "tiktok/kitex_gen/feed"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// List implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) List(ctx context.Context, req *feed.ListRequest) (resp *feed.ListResponse, err error) {
	// TODO: Your code here...
	return
}
