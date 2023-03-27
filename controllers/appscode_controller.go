/*
Copyright 2023.

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

package controllers

import (
	"context"
	"fmt"
	mycrdv1alpha1 "github.com/anisurrahman75/k8s-kubeBuilder/api/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	_ "k8s.io/client-go/informers/apps/v1"
	_ "k8s.io/client-go/listers/apps/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

var (
	deployOwnerKey = ".metadata.controller"
	svcOwnerKey    = ".metadata.controller"
	apiGVStr       = mycrdv1alpha1.GroupVersion.String()
	ourKind        = "CustomCrd"
)

// AppsCodeReconciler reconciles a AppsCode object
type AppsCodeReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=mycrd.k8s,resources=appscodes,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=mycrd.k8s,resources=appscodes/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=mycrd.k8s,resources=appscodes/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the AppsCode object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *AppsCodeReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)
	log.WithValues("ReqName", req.Name, "ReqNamespace", req.Namespace)

	// TODO(user): your logic here

	/*
		### 1: Load the CustomCrd by name
		We'll fetch the CustomCrd using our client.  All client methods take a
		context (to allow for cancellation) as their first argument, and the object
		in question as their last.  Get is a bit special, in that it takes a
		[`NamespacedName`](https://pkg.go.dev/sigs.k8s.io/controller-runtime/pkg/client?tab=doc#ObjectKey)
		as the middle argument (most don't have a middle argument, as we'll see
		below).
		Many client methods also take variadic options at the end.
	*/

	// appsCodeInstance have all data of AppsCode Resources
	var appsCodeInstance mycrdv1alpha1.AppsCode
	if err := r.Get(ctx, req.NamespacedName, &appsCodeInstance); err != nil {
		log.Error(err, "Unable to Fetch appsCodeInstance")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	var deploymentsInstance appsv1.Deployment

	if err := r.Client.Get(ctx, req.NamespacedName, &deploymentsInstance); err != nil {
		log.Info("could not find existing Deployment for ", appsCodeInstance.Name, ", creating one...")
		if errors.IsNotFound(err) {
			fmt.Println("New Deployment Creating... +++++++")
			if err = r.Client.Create(ctx, newDeployment(appsCodeInstance)); err != nil {
				log.Error(err, "Error Creating  %s Deployments", appsCodeInstance.Name)
			} else {
				fmt.Printf("%s Deployments Created...++++++\n", appsCodeInstance.Name)
			}
		}
	} else {
		expectedReplicas := int32(1)
		if appsCodeInstance.Spec.Replicas != nil {
			expectedReplicas = *appsCodeInstance.Spec.Replicas
		}
		if *deploymentsInstance.Spec.Replicas != expectedReplicas {
			log.Info("updating replica count", "old_count", *deploymentsInstance.Spec.Replicas, "new_count", expectedReplicas)
			deploymentsInstance.Spec.Replicas = &expectedReplicas
			fmt.Println(*deploymentsInstance.Spec.Replicas)
			if err := r.Client.Update(ctx, &deploymentsInstance); err != nil {
				fmt.Println("Failed")
				log.Error(err, "failed to Deployment update replica count")
				return ctrl.Result{}, err
			}
		}
	}

	var serviceInstance corev1.Service
	if err := r.Client.Get(ctx, req.NamespacedName, &serviceInstance); err != nil {
		log.Info("could not find existing service for ", appsCodeInstance.Name, ", creating one...")

		if errors.IsNotFound(err) {
			fmt.Println("New Service Creating... +++++++")
			if err = r.Client.Create(ctx, newService(appsCodeInstance)); err != nil {
				fmt.Println(err, "Error Creating ", appsCodeInstance.Name, "Service")
			} else {
				fmt.Printf("%s Services Created...+++++\n", appsCodeInstance.Name)
			}
		}
	}
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *AppsCodeReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&mycrdv1alpha1.AppsCode{}).
		Owns(&appsv1.Deployment{}).
		Owns(&corev1.Service{}).
		Complete(r)
}
