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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var helloapplog = logf.Log.WithName("helloapp-resource")

func (r *HelloApp) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-hp-mavrodiev-v1alpha1-helloapp,mutating=true,failurePolicy=fail,sideEffects=None,groups=hp.mavrodiev,resources=helloapps,verbs=create;update,versions=v1alpha1,name=mhelloapp.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &HelloApp{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *HelloApp) Default() {
	helloapplog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-hp-mavrodiev-v1alpha1-helloapp,mutating=false,failurePolicy=fail,sideEffects=None,groups=hp.mavrodiev,resources=helloapps,verbs=create;update,versions=v1alpha1,name=vhelloapp.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &HelloApp{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *HelloApp) ValidateCreate() error {
	helloapplog.Info("validate create", "name", r.Name)

	return r.ValidateHelloAppReplicaNumber()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *HelloApp) ValidateUpdate(old runtime.Object) error {
	helloapplog.Info("validate update", "name", r.Name)

	return r.ValidateHelloAppReplicaNumber()
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *HelloApp) ValidateDelete() error {
	helloapplog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}

func (r *HelloApp) ValidateHelloAppReplicaNumber() *field.Error {
	if r.Spec.Replicas < 1 || r.Spec.Replicas > 3 {
		return field.Invalid(field.NewPath("Spec").Child("Replicas"), r.Spec.Replicas, "Number of Replicas for this object should be in 1-3 limit")
	}
	return nil
}
