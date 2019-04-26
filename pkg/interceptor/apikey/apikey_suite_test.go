/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package apikey

import (
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"testing"
)

func TestApiKeyInterceptorPackage(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "API Key Interceptor package suite")
}
