package controllers

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ConsistOf", func() {
	It("should equal", func() {
		dep1 := createDeployment("nginx:latest")

		Expect(dep1).Should(HaveField("Namespace", Equal("default")))
		Expect(dep1).Should(HaveField("Name", Equal("nginx-deployment")))
		Expect(dep1).Should(HaveField("Labels", HaveKeyWithValue("app", "nginx")))
		Expect(dep1).Should(HaveField("Spec.Replicas", HaveValue(BeNumerically("==", 3))))
		Expect(dep1).ShouldNot(HaveField("Spec.Selector", BeNil()))
		Expect(dep1).Should(HaveField("Spec.Selector.MatchLabels", HaveKeyWithValue("app", "nginx")))
		Expect(dep1).Should(HaveField("Spec.Template.Labels", HaveKeyWithValue("app", "nginx")))
		Expect(dep1).ShouldNot(HaveField("Spec.Template.Spec.Containers", BeEmpty()))
		Expect(dep1.Spec.Template.Spec.Containers).Should(ConsistOf(SatisfyAll(
			HaveField("Name", Equal("nginx")),
			HaveField("Image", Equal("nginx:latest")),
			HaveField("Ports", Not(BeEmpty())),
			HaveField("Ports", ConsistOf(HaveField("ContainerPort", BeNumerically("==", 80)))),
		)))
	})
})
