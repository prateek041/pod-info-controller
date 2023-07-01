/*
Copyright 2023 prateek.io.

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

package controller

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	corev1 "k8s.io/api/core/v1"
	// testv1 "github.com/prateek041/controller-test.git/api/v1"
)

// PodReplicatorReconciler reconciles a PodReplicator object
type PodReplicatorReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=test.prateeksingh.tech,resources=podreplicators,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=test.prateeksingh.tech,resources=podreplicators/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=test.prateeksingh.tech,resources=podreplicators/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the PodReplicator object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.4/pkg/reconcile
func (r *PodReplicatorReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here
	pod := &corev1.Pod{}
	err := r.Get(ctx, req.NamespacedName, pod)

	if err != nil {
		log.Log.Error(err, "unable to fetch pod")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if pod.Status.Phase == corev1.PodRunning {
		log.Log.Info("pod fetched")
		fmt.Printf("\n Pod Name: %s", pod.Name)
		fmt.Printf("\n Pod Namespace: %s", pod.Namespace)
		fmt.Printf("\n Pod Labels: %s", pod.Spec.Containers[0].Image)
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *PodReplicatorReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		// for pod
		For(&corev1.Pod{}).
		Complete(r)
}
