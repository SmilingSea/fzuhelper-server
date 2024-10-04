// Code generated by Kitex v0.11.3. DO NOT EDIT.

package classroomservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	classroom "github.com/west2-online/fzuhelper-server/kitex_gen/classroom"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"GetEmptyRoom": kitex.NewMethodInfo(
		getEmptyRoomHandler,
		newClassroomServiceGetEmptyRoomArgs,
		newClassroomServiceGetEmptyRoomResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	classroomServiceServiceInfo                = NewServiceInfo()
	classroomServiceServiceInfoForClient       = NewServiceInfoForClient()
	classroomServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return classroomServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return classroomServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return classroomServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "ClassroomService"
	handlerType := (*classroom.ClassroomService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "classroom",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.11.3",
		Extra:           extra,
	}
	return svcInfo
}

func getEmptyRoomHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*classroom.ClassroomServiceGetEmptyRoomArgs)
	realResult := result.(*classroom.ClassroomServiceGetEmptyRoomResult)
	success, err := handler.(classroom.ClassroomService).GetEmptyRoom(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newClassroomServiceGetEmptyRoomArgs() interface{} {
	return classroom.NewClassroomServiceGetEmptyRoomArgs()
}

func newClassroomServiceGetEmptyRoomResult() interface{} {
	return classroom.NewClassroomServiceGetEmptyRoomResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetEmptyRoom(ctx context.Context, req *classroom.EmptyRoomRequest) (r *classroom.EmptyRoomResponse, err error) {
	var _args classroom.ClassroomServiceGetEmptyRoomArgs
	_args.Req = req
	var _result classroom.ClassroomServiceGetEmptyRoomResult
	if err = p.c.Call(ctx, "GetEmptyRoom", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}