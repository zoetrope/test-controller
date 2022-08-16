package v1

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"
)

var _ = Describe("Webhook Table Test", func() {

	DescribeTable("Validator Test", func(image string, replicas *int32, message string) {
		ctx := context.Background()
		sample := &Sample{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "default",
				Name:      "invalid-sample",
			},
			Spec: SampleSpec{
				Image:    image,
				Replicas: replicas,
			},
		}
		err := k8sClient.Create(ctx, sample)
		Expect(err).Should(HaveOccurred())

		Expect(err.Error()).Should(ContainSubstring(message))
	},
		Entry("replicas is negative", "nginx:1.14.2", pointer.Int32(-5), "replicas should be grater than 0"),
		Entry("image is empty", "", pointer.Int32(1), "image cannot be empty"),
		Entry("image does not have a tag", "nginx", pointer.Int32(1), "image should have a tag"),
		Entry("image is invalid format", "nginx:invalid:latest", pointer.Int32(1), "image is not valid format"),
		Entry("image has latest tag", "nginx:latest", pointer.Int32(1), "image cannot have latest tag"),
	)

})
