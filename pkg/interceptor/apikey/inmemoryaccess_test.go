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
	"fmt"
	"github.com/nalej/authx-interceptors/pkg/interceptor/config"
	"github.com/nalej/grpc-ping-go"
	"github.com/nalej/grpc-utils/pkg/test"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"math/rand"
)

var _ = ginkgo.Describe("InMemory Interceptor", func() {
	var targetServer *grpc.Server
	var targetListener *bufconn.Listener
	var pingClient grpc_ping_go.PingClient
	var validToken string
	var tokenProvider APIKeyAccess

	cfg := config.NewConfig(&config.AuthorizationConfig{
		AllowsAll: false,
		Permissions: map[string]config.Permission{
			"/ping.Ping/Ping": {Must: []string{KeyAccessPrimitive}},
		}}, "globalSecret", "authorization")

	targetListener = bufconn.Listen(test.BufSize)
	tokenProvider = NewInMemoryAPIKeyAccess()

	targetServer = grpc.NewServer(WithAPIKeyInterceptor(tokenProvider, cfg))

	// Launch Ping Server
	pingHandler := &PingHandler{}
	grpc_ping_go.RegisterPingServer(targetServer, pingHandler)
	test.LaunchServer(targetServer, targetListener)

	pingConn, err := test.GetConn(*targetListener)
	if err != nil {
		ginkgo.Fail("cannot obtain connection " + err.Error())
	}
	pingClient = grpc_ping_go.NewPingClient(pingConn)

	ginkgo.BeforeEach(func() {
		validToken = fmt.Sprintf("token_%d", rand.Intn(200))
		tokenProvider.(*InMemoryAPIKeyAccess).Clear()
		tokenProvider.(*InMemoryAPIKeyAccess).Add(validToken)
	})

	ginkgo.It("should be able to execute the command with a valid token", func() {
		ctx := getContext(cfg.Header, validToken)
		request := &grpc_ping_go.PingRequest{
			RequestNumber: 1,
		}
		response, err := pingClient.Ping(ctx, request)
		gomega.Expect(err).To(gomega.Succeed())
		gomega.Expect(response.RequestNumber).Should(gomega.Equal(request.RequestNumber))
	})

	ginkgo.It("should fail on an invalid token", func() {
		ctx := getContext(cfg.Header, "invalidToken")
		request := &grpc_ping_go.PingRequest{
			RequestNumber: 1,
		}
		_, err := pingClient.Ping(ctx, request)
		gomega.Expect(err).To(gomega.HaveOccurred())
	})
})
