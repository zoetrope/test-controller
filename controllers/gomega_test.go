package controllers

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("GomegaMatcher", func() {
	It("should equal", func() {
		dep1 := createDeployment("nginx:latest")

		Expect(dep1.Namespace).Should(Equal("default"))
		Expect(dep1.Name).Should(Equal("nginx-deployment"))
		Expect(dep1.Labels).Should(HaveKeyWithValue("app", "nginx"))
		Expect(dep1.Spec.Replicas).Should(HaveValue(BeNumerically("==", 3)))
		Expect(dep1.Spec.Selector).ShouldNot(BeNil())
		Expect(dep1.Spec.Selector.MatchLabels).Should(HaveKeyWithValue("app", "nginx"))
		Expect(dep1.Spec.Template.Labels).Should(HaveKeyWithValue("app", "nginx"))
		Expect(dep1.Spec.Template.Spec.Containers).ShouldNot(BeEmpty())
		Expect(dep1.Spec.Template.Spec.Containers[0].Name).Should(Equal("nginx"))
		Expect(dep1.Spec.Template.Spec.Containers[0].Image).Should(Equal("nginx:latest"))
		Expect(dep1.Spec.Template.Spec.Containers[0].Ports).ShouldNot(BeEmpty())
		Expect(dep1.Spec.Template.Spec.Containers[0].Ports[0].ContainerPort).Should(BeNumerically("==", 80))
	})
})
