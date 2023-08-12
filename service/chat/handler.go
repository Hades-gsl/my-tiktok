package main

import (
	"context"
	chat "tiktok/kitex_gen/chat"
)

// ChatServiceImpl implements the last service interface defined in the IDL.
type ChatServiceImpl struct{}

// ChatMessage implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) ChatMessage(ctx context.Context, req *chat.MessageRequest) (resp *chat.MessageResponse, err error) {
	// TODO: Your code here...
	return
}

// Send implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) Send(ctx context.Context, req *chat.SendRequest) (resp *chat.SendResponse, err error) {
	// TODO: Your code here...
	return
}
