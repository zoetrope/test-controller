package controllers

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("SemanticEqual", func() {
	ctx := context.TODO()

	It("should be equal", func() {
		dep1 := createDeployment("nginx:latest")
		dep2 := createDeployment("nginx:latest")
		dep1.Spec.Template.Spec.Containers[0].Resources = corev1.ResourceRequirements{
			Requests: corev1.ResourceList{
				corev1.ResourceCPU: resource.MustParse("1"),
			},
		}
		dep2.Spec.Template.Spec.Containers[0].Resources = corev1.ResourceRequirements{
			Requests: corev1.ResourceList{
				corev1.ResourceCPU: resource.MustParse("1000m"),
			},
		}

		Expect(dep1).ShouldNot(Equal(dep2))
		Expect(dep1).Should(SemanticEqual(dep2))
	})

	It("should equal original pod", func() {
		pod := &corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "default",
				Name:      "ubuntu",
			},
			Spec: corev1.PodSpec{
				Containers: []corev1.Container{
					{
						Name:  "ubuntu",
						Image: "ubuntu:20.04",
					},
				},
			},
		}
		originalPod := pod.DeepCopy()

		err := k8sClient.Create(ctx, pod)
		Expect(err).ShouldNot(HaveOccurred())

		Expect(pod.Spec).ShouldNot(Equal(originalPod.Spec))
		Expect(pod.Spec).Should(SemanticDerivative(originalPod.Spec))
	})
})
