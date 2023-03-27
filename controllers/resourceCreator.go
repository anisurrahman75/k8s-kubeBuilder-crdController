package controllers

import (
	mycrdv1alpha1 "github.com/anisurrahman75/k8s-kubeBuilder/api/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func newDeployment(appsCode mycrdv1alpha1.AppsCode) client.Object {
	labels := map[string]string{
		"app": appsCode.Spec.Name,
	}
	return &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind: "Deployment",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      appsCode.Name,
			Namespace: appsCode.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(&appsCode, mycrdv1alpha1.GroupVersion.WithKind(appsCode.Kind)),
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: appsCode.Spec.Replicas,
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
							Name:  appsCode.Spec.Name,
							Image: appsCode.Spec.Image,
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									Protocol:      corev1.ProtocolTCP,
									ContainerPort: appsCode.Spec.Port,
								},
							},
						},
					},
				},
			},
		},
	}
}
func newService(appsCode mycrdv1alpha1.AppsCode) client.Object {
	labels := map[string]string{
		"app": appsCode.Spec.Name,
	}
	return &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind: "Service",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      appsCode.Name,
			Namespace: appsCode.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(&appsCode, mycrdv1alpha1.GroupVersion.WithKind(appsCode.Kind)),
			},
		},
		Spec: corev1.ServiceSpec{
			Selector: labels,
			Type:     corev1.ServiceTypeNodePort,
			Ports: []corev1.ServicePort{
				{
					Name:       "apiserver",
					Port:       appsCode.Spec.Port,
					TargetPort: intstr.FromInt(int(appsCode.Spec.Port)),
					Protocol:   corev1.ProtocolTCP,
				},
			},
		},
	}
}
