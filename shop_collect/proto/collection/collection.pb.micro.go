// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/collection/collection.proto

package collection

import (
	fmt "fmt"
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

// Api Endpoints for Collection service

func NewCollectionEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Collection service

type CollectionService interface {
	CollectionList(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListsResponse, error)
	CollectionAdd(ctx context.Context, in *AddRequest, opts ...client.CallOption) (*MsgResponse, error)
	CollectionDel(ctx context.Context, in *DelRequest, opts ...client.CallOption) (*MsgResponse, error)
}

type collectionService struct {
	c    client.Client
	name string
}

func NewCollectionService(name string, c client.Client) CollectionService {
	return &collectionService{
		c:    c,
		name: name,
	}
}

func (c *collectionService) CollectionList(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListsResponse, error) {
	req := c.c.NewRequest(c.name, "Collection.CollectionList", in)
	out := new(ListsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionService) CollectionAdd(ctx context.Context, in *AddRequest, opts ...client.CallOption) (*MsgResponse, error) {
	req := c.c.NewRequest(c.name, "Collection.CollectionAdd", in)
	out := new(MsgResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionService) CollectionDel(ctx context.Context, in *DelRequest, opts ...client.CallOption) (*MsgResponse, error) {
	req := c.c.NewRequest(c.name, "Collection.CollectionDel", in)
	out := new(MsgResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Collection service

type CollectionHandler interface {
	CollectionList(context.Context, *ListRequest, *ListsResponse) error
	CollectionAdd(context.Context, *AddRequest, *MsgResponse) error
	CollectionDel(context.Context, *DelRequest, *MsgResponse) error
}

func RegisterCollectionHandler(s server.Server, hdlr CollectionHandler, opts ...server.HandlerOption) error {
	type collection interface {
		CollectionList(ctx context.Context, in *ListRequest, out *ListsResponse) error
		CollectionAdd(ctx context.Context, in *AddRequest, out *MsgResponse) error
		CollectionDel(ctx context.Context, in *DelRequest, out *MsgResponse) error
	}
	type Collection struct {
		collection
	}
	h := &collectionHandler{hdlr}
	return s.Handle(s.NewHandler(&Collection{h}, opts...))
}

type collectionHandler struct {
	CollectionHandler
}

func (h *collectionHandler) CollectionList(ctx context.Context, in *ListRequest, out *ListsResponse) error {
	return h.CollectionHandler.CollectionList(ctx, in, out)
}

func (h *collectionHandler) CollectionAdd(ctx context.Context, in *AddRequest, out *MsgResponse) error {
	return h.CollectionHandler.CollectionAdd(ctx, in, out)
}

func (h *collectionHandler) CollectionDel(ctx context.Context, in *DelRequest, out *MsgResponse) error {
	return h.CollectionHandler.CollectionDel(ctx, in, out)
}
