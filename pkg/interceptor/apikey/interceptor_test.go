/*
 * Copyright 2019 Nalej
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
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
func getContext(header string, token string) context.Context {
	md := metadata.New(map[string]string{header: token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	return ctx
}
