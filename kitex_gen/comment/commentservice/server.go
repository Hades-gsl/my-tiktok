// Code generated by Kitex v0.6.2. DO NOT EDIT.
package commentservice

import (
	server "github.com/cloudwego/kitex/server"
	comment "tiktok/kitex_gen/comment"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler comment.CommentService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
