/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package apikey

import (
	"context"
	"github.com/nalej/grpc-ping-go"
	"google.golang.org/grpc/metadata"
)

// PingHandler that receives device pings and will be launched with a device interceptor.
type PingHandler struct {
}

func (ph *PingHandler) Ping(ctx context.Context, request *grpc_ping_go.PingRequest) (*grpc_ping_go.PingResponse, error) {
	return &grpc_ping_go.PingResponse{
		RequestNumber: request.RequestNumber,
	}, nil
}

// GetContext returns a context with a given token.
func getContext(header string, token string) context.Context{
	md := metadata.New(map[string]string{header: token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	return ctx
}