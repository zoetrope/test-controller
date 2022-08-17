package controllers

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("Builder", func() {
	It("should equal", func() {
		dep1 := createDeployment("nginx:latest")

		dep2 := newDeployment("default", "nginx-deployment").
			withLabels(map[string]string{"app": "nginx"}).
			withReplicas(3).
			withNginxContainer("nginx:latest").
			build()

		Expect(dep2).Should(SemanticEqual(dep1))

		dep3 := newDeployment("default", "nginx-deployment").
			withLabels(map[string]string{"app": "nginx"}).
			withReplicas(3).
			withNginxContainer("nginx:latest").
			withSidecarContainer("ubuntu:22.04").
			build()

		Expect(dep3).ShouldNot(SemanticEqual(dep1))
	})
})

type deploymentBuilder struct {
	object *appsv1.Deployment
}

func newDeployment(namespace, name string) *deploymentBuilder {
	return &deploymentBuilder{
		object: &appsv1.Deployment{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: namespace,
				Name:      name,
			},
		},
	}
}

func (b *deploymentBuilder) withLabels(labels map[string]string) *deploymentBuilder {
	if b.object.Labels == nil {
		b.object.Labels = map[string]string{}
	}

	for key, value := range labels {
		b.object.Labels[key] = value
	}
	return b
}

func (b *deploymentBuilder) withReplicas(replicas int32) *deploymentBuilder {
	b.object.Spec.Replicas = &replicas
	return b
}

func (b *deploymentBuilder) withNginxContainer(image string) *deploymentBuilder {
	b.object.Spec.Selector = &metav1.LabelSelector{
		MatchLabels: map[string]string{
			"app": "nginx",
		},
	}
	b.object.Spec.Template.Labels = map[string]string{
		"app": "nginx",
	}
	b.object.Spec.Template.Spec.Containers =
		append(b.object.Spec.Template.Spec.Containers,
			corev1.Container{
				Name:  "nginx",
				Image: image,
				Ports: []corev1.ContainerPort{
					{
						ContainerPort: 80,
					},
				},
			})
	return b
}

func (b *deploymentBuilder) withSidecarContainer(image string) *deploymentBuilder {
	b.object.Spec.Template.Spec.Containers =
		append(b.object.Spec.Template.Spec.Containers,
			corev1.Container{
				Name:  "sidecar",
				Image: image,
			})
	return b
}

func (b *deploymentBuilder) build() *appsv1.Deployment {
	return b.object
}
