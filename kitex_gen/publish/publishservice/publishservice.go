// Code generated by Kitex v0.7.0. DO NOT EDIT.

package publishservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	publish "tiktok/kitex_gen/publish"
)

func serviceInfo() *kitex.ServiceInfo {
	return publishServiceServiceInfo
}

var publishServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "PublishService"
	handlerType := (*publish.PublishService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Action": kitex.NewMethodInfo(actionHandler, newPublishServiceActionArgs, newPublishServiceActionResult, false),
		"List":   kitex.NewMethodInfo(listHandler, newPublishServiceListArgs, newPublishServiceListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "publish",
		"ServiceFilePath": "idl/publish.thrift",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.7.0",
		Extra:           extra,
	}
	return svcInfo
}

func actionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*publish.PublishServiceActionArgs)
	realResult := result.(*publish.PublishServiceActionResult)
	success, err := handler.(publish.PublishService).Action(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPublishServiceActionArgs() interface{} {
	return publish.NewPublishServiceActionArgs()
}

func newPublishServiceActionResult() interface{} {
	return publish.NewPublishServiceActionResult()
}

func listHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*publish.PublishServiceListArgs)
	realResult := result.(*publish.PublishServiceListResult)
	success, err := handler.(publish.PublishService).List(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPublishServiceListArgs() interface{} {
	return publish.NewPublishServiceListArgs()
}

func newPublishServiceListResult() interface{} {
	return publish.NewPublishServiceListResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Action(ctx context.Context, req *publish.ActionRequest) (r *publish.ActionResponse, err error) {
	var _args publish.PublishServiceActionArgs
	_args.Req = req
	var _result publish.PublishServiceActionResult
	if err = p.c.Call(ctx, "Action", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) List(ctx context.Context, req *publish.ListRequest) (r *publish.ListResponse, err error) {
	var _args publish.PublishServiceListArgs
	_args.Req = req
	var _result publish.PublishServiceListResult
	if err = p.c.Call(ctx, "List", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
