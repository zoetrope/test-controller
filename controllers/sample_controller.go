package controllers

import (
	"context"

	testv1 "github.com/zoetrope/test-controller/api/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// SampleReconciler reconciles a Sample object
type SampleReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=test.zoetrope.github.io,resources=samples,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=test.zoetrope.github.io,resources=samples/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=test.zoetrope.github.io,resources=samples/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *SampleReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	var sample testv1.Sample
	err := r.Get(ctx, req.NamespacedName, &sample)
	if err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	if !sample.DeletionTimestamp.IsZero() {
		return ctrl.Result{}, nil
	}

	dep := &appsv1.Deployment{}
	dep.SetNamespace(sample.Namespace)
	dep.SetName(sample.Name)

	op, err := ctrl.CreateOrUpdate(ctx, r.Client, dep, func() error {
		if dep.Labels == nil {
			dep.Labels = map[string]string{}
		}
		dep.Labels["app.kubernetes.io/name"] = "nginx"
		dep.Labels["app.kubernetes.io/instance"] = req.Name

		dep.Spec = appsv1.DeploymentSpec{
			Replicas: sample.Spec.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app.kubernetes.io/name":     "nginx",
					"app.kubernetes.io/instance": req.Name,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app.kubernetes.io/name":     "nginx",
						"app.kubernetes.io/instance": req.Name,
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "nginx",
							Image: sample.Spec.Image,
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		}
		return ctrl.SetControllerReference(&sample, dep, r.Scheme)
	})
	if err != nil {
		logger.Error(err, "failed to create or update Deployment")
		return ctrl.Result{}, err
	}

	if op != controllerutil.OperationResultNone {
		logger.Info("reconcile Deployment successfully", "op", op)
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *SampleReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&testv1.Sample{}).
		Owns(&appsv1.Deployment{}).
		Complete(r)
}
