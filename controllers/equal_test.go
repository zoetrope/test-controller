package controllers

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Equal", func() {
	It("should equal", func() {
		dep1 := createDeployment("nginx:latest")
		dep2 := createDeployment("nginx:latest")

		Expect(dep1).Should(Equal(dep2))
	})
})
