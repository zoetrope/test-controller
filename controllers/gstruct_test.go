package controllers

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

var _ = Describe("gstruct", func() {
	It("should equal", func() {
		dep1 := createDeployment("nginx:latest")

		Expect(dep1).Should(PointTo(MatchFields(IgnoreExtras, Fields{
			"ObjectMeta": MatchFields(IgnoreExtras, Fields{
				"Namespace": Equal("default"),
				"Name":      Equal("nginx-deployment"),
				"Labels":    MatchAllKeys(Keys{"app": Equal("nginx")}),
			}),
			"Spec": MatchFields(IgnoreExtras, Fields{
				"Replicas": PointTo(BeNumerically("==", 3)),
				"Selector": PointTo(MatchFields(IgnoreExtras, Fields{
					"MatchLabels": MatchAllKeys(Keys{"app": Equal("nginx")}),
				})),
				"Template": MatchFields(IgnoreExtras, Fields{
					"ObjectMeta": MatchFields(IgnoreExtras, Fields{
						"Labels": MatchAllKeys(Keys{"app": Equal("nginx")}),
					}),
					"Spec": MatchFields(IgnoreExtras, Fields{
						"Containers": MatchAllElements(containerIdentity, Elements{
							"nginx": MatchFields(IgnoreExtras, Fields{
								"Image": Equal("nginx:latest"),
								"Ports": MatchAllElements(portIdentity, Elements{
									"80": HaveField("ContainerPort", BeNumerically("==", 80)),
								}),
							}),
						},
						),
					}),
				}),
			}),
		})))
	})
})
