// SPDX-FileCopyrightText: Contributors to the Gardener project
//
// SPDX-License-Identifier: Apache-2.0

package etcdmember

import (
	"context"
	"fmt"
	"slices"
	"time"

	druidcorev1alpha1 "github.com/gardener/etcd-druid/api/core/v1alpha1"
	machinev1alpha1 "github.com/gardener/machine-controller-manager/pkg/apis/machine/v1alpha1"
	"github.com/go-logr/logr"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/clock"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	v1beta1constants "github.com/gardener/gardener/pkg/apis/core/v1beta1/constants"
	"github.com/gardener/gardener/pkg/component/etcd/etcd"
	"github.com/gardener/gardener/pkg/extensions"
	"github.com/gardener/gardener/pkg/utils/flow"
	gardenerutils "github.com/gardener/gardener/pkg/utils/gardener"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
)

// Reconciler ensures the control plane machine's address is added to Etcd's `.spec.externallyManagedMemberAddresses[]`
// list. If instance deletion is suspended, it removes it from there.
type Reconciler struct {
	Client client.Client
	Clock  clock.Clock
}

// Reconcile ensures the control plane machine's address is added to Etcd's `.spec.externallyManagedMemberAddresses[]`
// list. If instance deletion is suspended, it removes it from there.
func (r *Reconciler) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	log := logf.FromContext(ctx)

	machine := &machinev1alpha1.Machine{}
	if err := r.Client.Get(ctx, request.NamespacedName, machine); err != nil {
		if apierrors.IsNotFound(err) {
			log.V(1).Info("Object is gone, stop reconciling")
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, fmt.Errorf("error retrieving object from store: %w", err)
	}

	address, found := gardenerutils.PreferredNodeAddress(machine.Status.Addresses)
	if !found {
		log.Info("No address found in machine's .status.addresses[] list, exiting early")
		return reconcile.Result{}, nil
	}
	log = log.WithValues("address", address)

	if machine.DeletionTimestamp == nil {
		log.Info("Ensure machine address is added to ETCD objects")
		return reconcile.Result{}, r.ensureMachineIPInBothETCDs(log, machine, address, addAddressToETCD)(ctx)
	}

	if instanceDeletionSuspended(machine) {
		log.Info("Ensure machine address is removed from ETCD objects")
		if err := r.ensureMachineIPInBothETCDs(log, machine, address, removeAddressFromETCD)(ctx); err != nil {
			return reconcile.Result{}, err
		}

		log.Info("Machine IP got removed from ETCD objects, now removing suspend annotation from Machine object", "annotationKey", gardenerutils.AnnotationKeySuspendInstanceDeletion)
		patch := client.MergeFrom(machine.DeepCopy())
		delete(machine.Annotations, gardenerutils.AnnotationKeySuspendInstanceDeletion)
		return reconcile.Result{}, r.Client.Patch(ctx, machine, patch)
	}

	return reconcile.Result{}, nil
}

func (r *Reconciler) ensureMachineIPInBothETCDs(log logr.Logger, machine *machinev1alpha1.Machine, address string, mutate func(*druidcorev1alpha1.Etcd, string) bool) flow.TaskFn {
	return flow.Parallel(
		func(ctx context.Context) error {
			return r.ensureMachineIPInETCD(ctx, log, machine, v1beta1constants.ETCDRoleMain, address, mutate)
		},
		func(ctx context.Context) error {
			return r.ensureMachineIPInETCD(ctx, log, machine, v1beta1constants.ETCDRoleEvents, address, mutate)
		},
	)
}

func (r *Reconciler) ensureMachineIPInETCD(ctx context.Context, log logr.Logger, machine *machinev1alpha1.Machine, role, address string, mutate func(*druidcorev1alpha1.Etcd, string) bool) error {
	etcdObj := &druidcorev1alpha1.Etcd{ObjectMeta: metav1.ObjectMeta{Name: etcd.Name(role), Namespace: machine.Namespace}}
	if err := r.Client.Get(ctx, client.ObjectKeyFromObject(etcdObj), etcdObj); err != nil {
		return fmt.Errorf("failed reading etcd %s: %w", client.ObjectKeyFromObject(etcdObj), err)
	}

	patch := client.MergeFromWithOptions(etcdObj.DeepCopy(), client.MergeFromWithOptimisticLock{})
	if mutate(etcdObj, address) {
		metav1.SetMetaDataAnnotation(&etcdObj.ObjectMeta, v1beta1constants.GardenerOperation, v1beta1constants.GardenerOperationReconcile)
		metav1.SetMetaDataAnnotation(&etcdObj.ObjectMeta, v1beta1constants.GardenerTimestamp, r.Clock.Now().UTC().Format(time.RFC3339Nano))
		if err := r.Client.Patch(ctx, etcdObj, patch); err != nil {
			return fmt.Errorf("failed patching prefered address %s for machine in etcd spec: %w", address, err)
		}
	}

	return extensions.WaitUntilObjectReadyWithHealthFunction(
		ctx,
		r.Client,
		log,
		etcd.CheckEtcdObject,
		etcdObj,
		"Etcd",
		etcd.DefaultInterval,
		etcd.DefaultSevereThreshold,
		etcd.DefaultTimeout,
		nil,
	)
}

func addAddressToETCD(etcdObj *druidcorev1alpha1.Etcd, address string) bool {
	if slices.Contains(etcdObj.Spec.ExternallyManagedMemberAddresses, address) {
		return false
	}
	etcdObj.Spec.ExternallyManagedMemberAddresses = append(etcdObj.Spec.ExternallyManagedMemberAddresses, address)
	etcdObj.Spec.Replicas++
	return true
}

func removeAddressFromETCD(etcdObj *druidcorev1alpha1.Etcd, address string) bool {
	if !slices.Contains(etcdObj.Spec.ExternallyManagedMemberAddresses, address) {
		return false
	}
	etcdObj.Spec.ExternallyManagedMemberAddresses = slices.DeleteFunc(etcdObj.Spec.ExternallyManagedMemberAddresses, func(v string) bool {
		return v == address
	})
	etcdObj.Spec.Replicas--
	return true
}
