package controllers

import (
	"context"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Debug mode", func() {
	ctx := context.TODO()

	AfterEach(func() {
		suiteConfig, _ := GinkgoConfiguration()
		if CurrentSpecReport().Failed() && suiteConfig.FailFast {
			suiteFailed = true
		}
	})

	It("should create deployment", func() {
		dep := createDeployment("nginx:latest")

		err := k8sClient.Create(ctx, dep)
		Expect(err).ShouldNot(HaveOccurred())
	})
})
