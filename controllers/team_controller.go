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

package controllers

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	orgv1alpha1 "github.com/jeremymv2/team-controller/api/v1alpha1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
)

// TeamReconciler reconciles a Team object
type TeamReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=org.ethzero.cloud,resources=teams,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=org.ethzero.cloud,resources=teams/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=org.ethzero.cloud,resources=teams/finalizers,verbs=update
//+kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=roles,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=rolebindings,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Team object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.0/pkg/reconcile
func (r *TeamReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)
	/*
		### 1: Load the Team by name
		We'll fetch the Team using our client.  All client methods take a
		context (to allow for cancellation) as their first argument, and the object
		in question as their last.  Get is a bit special, in that it takes a
		[`NamespacedName`](https://pkg.go.dev/sigs.k8s.io/controller-runtime/pkg/client?tab=doc#ObjectKey)
		as the middle argument (most don't have a middle argument, as we'll see
		below).
		Many client methods also take variadic options at the end.
	*/
	var team orgv1alpha1.Team

	if err := r.Get(ctx, req.NamespacedName, &team); err != nil {
		log.Error(err, "unable to fetch Team")
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	nsSpec := &v1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: team.Name}}

	if err := r.Create(ctx, nsSpec); err != nil {
		log.Error(err, "unable to create Namespace for Team", "team", team.Name)
		return ctrl.Result{}, err
	}

	log.V(1).Info("created Namespace for Team", "team", team.Name)

	/*
		roleBinding := &rbacv1.RoleBinding{
			TypeMeta: metav1.TypeMeta{
				Kind:       "RoleBinding",
				APIVersion: "rbac.authorization.k8s.io/v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      fmt.Sprintf("ad-kubernetes-%s", req.NamespacedName),
				Namespace: namespaceName,
			},
			Subjects: []v1beta1.Subject{
				v1beta1.Subject{
					Kind: "Group",
					Name: fmt.Sprintf("ad-kubernetes-%s", namespaceName),
				},
			},
			RoleRef: v1beta1.RoleRef{
				APIGroup: "rbac.authorization.k8s.io",
				Kind:     "ClusterRole",
				Name:     "edit",
			},
		}
	*/

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *TeamReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&orgv1alpha1.Team{}).
		Complete(r)
}
