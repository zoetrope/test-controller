package controllers

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	testv1 "github.com/zoetrope/test-controller/api/v1"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ = Describe("Test SampleController", func() {
	ctx := context.Background()
	var stopFunc func()

	BeforeEach(func() {
		mgr, err := ctrl.NewManager(cfg, ctrl.Options{
			Scheme:             scheme,
			LeaderElection:     false,
			MetricsBindAddress: "0",
		})
		Expect(err).ShouldNot(HaveOccurred())

		reconciler := &SampleReconciler{
			Client: mgr.GetClient(),
			Scheme: scheme,
		}
		err = reconciler.SetupWithManager(mgr)
		Expect(err).ShouldNot(HaveOccurred())

		ctx, cancel := context.WithCancel(ctx)
		stopFunc = cancel
		go func() {
			err := mgr.Start(ctx)
			if err != nil {
				panic(err)
			}
		}()
		time.Sleep(100 * time.Millisecond)
	})

	AfterEach(func() {
		stopFunc()
		time.Sleep(100 * time.Millisecond)
	})

	It("should be success", func() {
		sample := &testv1.Sample{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "default",
				Name:      "test",
			},
			Spec: testv1.SampleSpec{
				Image:    "nginx:1.14.2",
				Replicas: pointer.Int32(2),
			},
		}

		err := k8sClient.Create(ctx, sample)
		Expect(err).ShouldNot(HaveOccurred())

		Eventually(func(g Gomega) {
			dep := &appsv1.Deployment{}
			err = k8sClient.Get(ctx, client.ObjectKey{Namespace: "default", Name: "test"}, dep)
			g.Expect(err).ShouldNot(HaveOccurred())

			g.Expect(dep.Spec.Replicas).Should(HaveValue(BeNumerically("==", 2)))
			g.Expect(dep.Spec.Template.Spec.Containers).Should(MatchAllElements(containerIdentity, Elements{
				"nginx": HaveField("Image", "nginx:1.14.2"),
			}))
		}).Should(Succeed())
	})
})
