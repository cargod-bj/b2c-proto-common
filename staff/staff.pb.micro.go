// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: staff/staff.proto

package staff

import (
	fmt "fmt"
	common "github.com/cargod-bj/b2c-proto/common"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Staff service

func NewStaffEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Staff service

type StaffService interface {
	Login(ctx context.Context, in *Request, opts ...client.CallOption) (*common.Response, error)
	Register(ctx context.Context, in *User, opts ...client.CallOption) (*common.Response, error)
}

type staffService struct {
	c    client.Client
	name string
}

func NewStaffService(name string, c client.Client) StaffService {
	return &staffService{
		c:    c,
		name: name,
	}
}

func (c *staffService) Login(ctx context.Context, in *Request, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Staff.Login", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffService) Register(ctx context.Context, in *User, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Staff.Register", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Staff service

type StaffHandler interface {
	Login(context.Context, *Request, *common.Response) error
	Register(context.Context, *User, *common.Response) error
}

func RegisterStaffHandler(s server.Server, hdlr StaffHandler, opts ...server.HandlerOption) error {
	type staff interface {
		Login(ctx context.Context, in *Request, out *common.Response) error
		Register(ctx context.Context, in *User, out *common.Response) error
	}
	type Staff struct {
		staff
	}
	h := &staffHandler{hdlr}
	return s.Handle(s.NewHandler(&Staff{h}, opts...))
}

type staffHandler struct {
	StaffHandler
}

func (h *staffHandler) Login(ctx context.Context, in *Request, out *common.Response) error {
	return h.StaffHandler.Login(ctx, in, out)
}

func (h *staffHandler) Register(ctx context.Context, in *User, out *common.Response) error {
	return h.StaffHandler.Register(ctx, in, out)
}
