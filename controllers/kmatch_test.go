package controllers

import (
	. "github.com/kralicky/kmatch"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("kmatch", func() {
	It("should equal", func() {
		dep1 := createDeployment("nginx:latest")

		Expect(dep1).Should(SatisfyAll(
			HaveNamespace("default"),
			HaveName("nginx-deployment"),
			HaveLabels("app", "nginx"),
			HaveReplicaCount(3),
			HaveMatchingContainer(SatisfyAll(
				HaveName("nginx"),
				HaveImage("nginx:latest"),
				HavePorts(80),
			)),
		))
	})
})
