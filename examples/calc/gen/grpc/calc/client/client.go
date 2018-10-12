// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// calc GRPC client
//
// Command:
// $ goa gen goa.design/goa/examples/calc/design -o
// $(GOPATH)/src/goa.design/goa/examples/calc

package client

import (
	"context"

	goa "goa.design/goa"
	goagrpc "goa.design/goa/grpc"
)

// Client lists the service endpoint gRPC clients.
type Client struct {
}

// NewClient instantiates gRPC client for all the calc service servers.
func NewClient() *Client {
	return &Client{}
}

// Add calls the "Add" function in calcpb.CalcClient interface.
func (c *Client) Add(doer goagrpc.UnaryDoer) goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		return doer.Invoke(ctx, v)
	}
}
