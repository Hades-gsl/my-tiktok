package main

import (
	"context"
	publish "tiktok/kitex_gen/publish"
)

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

// Action implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) Action(ctx context.Context, req *publish.ActionRequest) (resp *publish.ActionResponse, err error) {
	// TODO: Your code here...
	return
}

// List implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) List(ctx context.Context, req *publish.ListRequest) (resp *publish.ListResponse, err error) {
	// TODO: Your code here...
	return
}
