package main

import (
	"context"
	gateway "zhamghaoran/ddbr-gateway/kitex_gen/ddbr/rpc/gateway"
)

// GatewayImpl implements the last service interface defined in the IDL.
type GatewayImpl struct{}

// Set implements the GatewayImpl interface.
func (s *GatewayImpl) Set(ctx context.Context, req *gateway.SetRequest) (resp *gateway.SetResponse, err error) {
	// TODO: Your code here...
	return
}

// Get implements the GatewayImpl interface.
func (s *GatewayImpl) Get(ctx context.Context, req *gateway.GetRequest) (resp *gateway.GetResponse, err error) {
	resp = gateway.NewGetResponse()
	resp.Val = "pong"
	return
}

// RegisterSever implements the GatewayImpl interface.
func (s *GatewayImpl) RegisterSever(ctx context.Context, req *gateway.RegisterSeverReq) (resp *gateway.RegisterSeverResp, err error) {
	// TODO: Your code here...
	return
}

// RegisterGateway implements the GatewayImpl interface.
func (s *GatewayImpl) RegisterGateway(ctx context.Context, req *gateway.RegisterGatewayReq) (resp *gateway.RegisterGatewayResp, err error) {
	// TODO: Your code here...
	return
}
