/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	"strings"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var samplelog = logf.Log.WithName("sample-resource")

func (r *Sample) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-test-zoetrope-github-io-v1-sample,mutating=true,failurePolicy=fail,sideEffects=None,groups=test.zoetrope.github.io,resources=samples,verbs=create;update,versions=v1,name=msample.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &Sample{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *Sample) Default() {
	samplelog.Info("default", "name", r.Name)

	defaultReplicas := int32(1)

	if r.Spec.Replicas == nil {
		r.Spec.Replicas = &defaultReplicas
	}
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-test-zoetrope-github-io-v1-sample,mutating=false,failurePolicy=fail,sideEffects=None,groups=test.zoetrope.github.io,resources=samples,verbs=create;update,versions=v1,name=vsample.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &Sample{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Sample) ValidateCreate() error {
	samplelog.Info("validate create", "name", r.Name)

	return r.validate()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Sample) ValidateUpdate(old runtime.Object) error {
	samplelog.Info("validate update", "name", r.Name)

	return r.validate()
}

func (r *Sample) validate() error {
	var errs field.ErrorList

	if r.Spec.Replicas == nil {
		errs = append(errs, field.Invalid(field.NewPath("spec", "replicas"), r.Spec.Replicas, "replicas cannot be empty"))
	} else if *r.Spec.Replicas < 1 {
		errs = append(errs, field.Invalid(field.NewPath("spec", "replicas"), r.Spec.Replicas, "replicas should be grater than 0"))
	}

	if len(r.Spec.Image) == 0 {
		errs = append(errs, field.Invalid(field.NewPath("spec", "image"), r.Spec.Image, "image cannot be empty"))
	} else if !strings.Contains(r.Spec.Image, ":") {
		errs = append(errs, field.Invalid(field.NewPath("spec", "image"), r.Spec.Image, "image should have a tag"))
	} else {
		images := strings.Split(r.Spec.Image, ":")
		if len(images) != 2 {
			errs = append(errs, field.Invalid(field.NewPath("spec", "image"), r.Spec.Image, "image is not valid formt"))
		} else if images[1] == "latest" {
			errs = append(errs, field.Invalid(field.NewPath("spec", "image"), r.Spec.Image, "image cannot have latest tag"))
		}
	}

	if len(errs) > 0 {
		err := apierrors.NewInvalid(schema.GroupKind{Group: GroupVersion.Group, Kind: "Sample"}, r.Name, errs)
		return err
	}
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Sample) ValidateDelete() error {
	samplelog.Info("validate delete", "name", r.Name)

	return nil
}
