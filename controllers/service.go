package controllers

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	hpv1alpha1 "github.com/h-mavrodiev/hello-app-operator/api/v1alpha1"
)

// Ensures Service is Running in a Namespace
func (r HelloAppReconciler) ensureService(request reconcile.Request,
	instance *hpv1alpha1.HelloApp,
	service *corev1.Service,
) (*reconcile.Result, error) {

	// See if service exists. Create it if it doesn`t
	found := &corev1.Service{}
	err := r.Get(context.TODO(), types.NamespacedName{
		Name:      service.Name,
		Namespace: instance.Namespace,
	}, found)
	if err != nil && errors.IsNotFound(err) {

		// Create Service
		err = r.Create(context.TODO(), service)

		if err != nil {
			// Service creation failed
			return &reconcile.Result{}, err
		} else {
			// Service creation was successful
			return nil, nil
		}
	} else if err != nil {
		// Error that isn`t due to the service not existing
		return &reconcile.Result{}, err
	}

	return nil, nil
}

// Code for creating a Service
func (r *HelloAppReconciler) backendService(v *hpv1alpha1.HelloApp) *corev1.Service {

	var labels map[string]string = map[string]string{"name": v.Name}

	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      v.Name,
			Namespace: v.Namespace,
		},
		Spec: corev1.ServiceSpec{
			Selector: labels,
			Ports: []corev1.ServicePort{{
				Protocol:   corev1.ProtocolTCP,
				Port:       80,
				TargetPort: intstr.FromInt(8080),
			}},
			// ServiceTypeClusterIP means a service will only be accessible inside the
			// cluster, via the cluster IP.
			Type: corev1.ServiceTypeClusterIP,
		},
	}

	controllerutil.SetControllerReference(v, service, r.Scheme)
	return service
}
