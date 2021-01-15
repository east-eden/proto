// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: combat/combat.proto

package combat

import (
	_ "e.coding.net/mmstudio/blade/proto/go_out/game"
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

// Api Endpoints for CombatService service

func NewCombatServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CombatService service

type CombatService interface {
	StartStageCombat(ctx context.Context, in *StartStageCombatReq, opts ...client.CallOption) (*StartStageCombatReply, error)
}

type combatService struct {
	c    client.Client
	name string
}

func NewCombatService(name string, c client.Client) CombatService {
	return &combatService{
		c:    c,
		name: name,
	}
}

func (c *combatService) StartStageCombat(ctx context.Context, in *StartStageCombatReq, opts ...client.CallOption) (*StartStageCombatReply, error) {
	req := c.c.NewRequest(c.name, "CombatService.StartStageCombat", in)
	out := new(StartStageCombatReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CombatService service

type CombatServiceHandler interface {
	StartStageCombat(context.Context, *StartStageCombatReq, *StartStageCombatReply) error
}

func RegisterCombatServiceHandler(s server.Server, hdlr CombatServiceHandler, opts ...server.HandlerOption) error {
	type combatService interface {
		StartStageCombat(ctx context.Context, in *StartStageCombatReq, out *StartStageCombatReply) error
	}
	type CombatService struct {
		combatService
	}
	h := &combatServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CombatService{h}, opts...))
}

type combatServiceHandler struct {
	CombatServiceHandler
}

func (h *combatServiceHandler) StartStageCombat(ctx context.Context, in *StartStageCombatReq, out *StartStageCombatReply) error {
	return h.CombatServiceHandler.StartStageCombat(ctx, in, out)
}
