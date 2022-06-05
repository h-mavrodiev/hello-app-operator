package controllers

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	hpv1alpha1 "github.com/h-mavrodiev/hello-app-operator/api/v1alpha1"
)

// Ensures Deployment resource presence in given Namespace
func (r *HelloAppReconciler) ensureDeployment(request reconcile.Request,
	instance *hpv1alpha1.HelloApp, deployment *appsv1.Deployment) (*reconcile.Result, error) {

	//See if Deployment already exists. If not create it.
	found := &appsv1.Deployment{}
	err := r.Get(context.TODO(), types.NamespacedName{
		Name:      deployment.Name,
		Namespace: instance.Namespace,
	}, found)
	if err != nil && errors.IsNotFound(err) {

		// Create the Deployment
		err = r.Create(context.TODO(), deployment)

		if err != nil {
			// Deployment failed
			return &reconcile.Result{}, err
		} else {
			// Deployment was successful
			return nil, nil
		}

	} else if err != nil {
		// Error that isn`t due to the Deployment not existing
		return &reconcile.Result{}, err
	}

	return nil, nil
}

// Code for creating Deployment
func (r *HelloAppReconciler) backendDeployment(v *hpv1alpha1.HelloApp) *appsv1.Deployment {

	var labels map[string]string = map[string]string{"name": v.Name}
	size := int32(1)

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      v.Name,
			Namespace: v.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &size,
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Image:           v.Spec.Image,
							ImagePullPolicy: corev1.PullAlways,
							Name:            "hello-app-pod",
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: 8080,
								},
							},
							LivenessProbe: &corev1.Probe{
								ProbeHandler: corev1.ProbeHandler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "/healthz",
										Port: intstr.IntOrString{
											IntVal: 8080,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	controllerutil.SetControllerReference(v, deployment, r.Scheme)
	return deployment
}
