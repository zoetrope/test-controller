package v1

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"
)

var _ = Describe("Webhook Test", func() {
	ctx := context.Background()

	It("should deny image that has latest tag", func() {
		sample := &Sample{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "default",
				Name:      "invalid-sample",
			},
			Spec: SampleSpec{
				Image:    "nginx:latest",
				Replicas: pointer.Int32(2),
			},
		}
		err := k8sClient.Create(ctx, sample)

		Expect(err).Should(HaveStatusErrorReason(Equal(metav1.StatusReasonInvalid)))
		Expect(err).Should(HaveStatusErrorMessage(ContainSubstring("image cannot have latest tag")))
	})
})
