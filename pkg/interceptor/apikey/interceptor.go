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
	"github.com/nalej/authx-interceptors/pkg/interceptor/config"
	"github.com/nalej/derrors"
	"github.com/nalej/grpc-utils/pkg/conversions"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// WithAPIKeyInterceptor is a gRPC option. If this option is included, the interceptor verifies that the API Key
// is authorized to use the method.
func WithAPIKeyInterceptor(apiKeyAccess APIKeyAccess, config *config.Config) grpc.ServerOption {
	return grpc.UnaryInterceptor(keyInterceptor(apiKeyAccess, config))
}

func keyInterceptor(apiKeyAccess APIKeyAccess, config *config.Config) grpc.UnaryServerInterceptor {
	return func(ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {

		_, ok := config.Authorization.Permissions[info.FullMethod]

		if ok {
			// Extract the API token.
			tk, err := extractToken(ctx, config)
			if err != nil {
				return nil, conversions.ToGRPCError(derrors.NewUnauthenticatedError("token is not supplied"))
			}
			// Check the claim using the extracted token.
			claim, dErr := checkToken(*tk, apiKeyAccess)
			if dErr != nil {
				return nil, conversions.ToGRPCError(dErr)
			}
			dErr = config.AuthorizePrimitive(info.FullMethod, claim.Primitives)

			if dErr != nil {
				return nil, conversions.ToGRPCError(dErr)
			}
			return handler(ctx, req)

		} else {
			if !config.Authorization.AllowsAll {
				return nil, conversions.ToGRPCError(
					derrors.NewUnauthenticatedError("unauthorized method").
						WithParams(info.FullMethod))
			}
		}
		log.Warn().Msg("auth metadata has not been added")
		return handler(ctx, req)
	}

}

// ExtractToken obtains the token from the Authorization header.
func extractToken(ctx context.Context, config *config.Config) (*string, derrors.Error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, derrors.NewInternalError("impossible to extract metadata")
	}
	authHeader, ok := md[config.Header]
	if !ok {
		return nil, derrors.NewUnauthenticatedError("token is not supplied")
	}
	rawToken := authHeader[0]
	return &rawToken, nil
}

// CheckToken checks that the token is valid
func checkToken(token string, keyAccess APIKeyAccess) (*KeyClaim, derrors.Error) {
	err := keyAccess.IsValid(token)
	if err != nil {
		return nil, err
	}
	return NewDefaultKeyClaim(), nil
}
