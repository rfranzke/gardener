// Copyright 2021 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package app

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"

	druidv1alpha1 "github.com/gardener/etcd-druid/api/v1alpha1"
	machinev1alpha1 "github.com/gardener/machine-controller-manager/pkg/apis/machine/v1alpha1"
	"github.com/spf13/cobra"
	admissionregistrationv1 "k8s.io/api/admissionregistration/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	vpaautoscalingv1 "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/apis/autoscaling.k8s.io/v1"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	"github.com/gardener/gardener/extensions/pkg/controller"
	extensionscmdcontroller "github.com/gardener/gardener/extensions/pkg/controller/cmd"
	"github.com/gardener/gardener/extensions/pkg/controller/controlplane/genericactuator"
	"github.com/gardener/gardener/extensions/pkg/controller/heartbeat"
	extensionsheartbeatcmd "github.com/gardener/gardener/extensions/pkg/controller/heartbeat/cmd"
	"github.com/gardener/gardener/extensions/pkg/controller/operatingsystemconfig/oscommon"
	"github.com/gardener/gardener/extensions/pkg/controller/worker"
	extensionscmdwebhook "github.com/gardener/gardener/extensions/pkg/webhook/cmd"
	v1beta1constants "github.com/gardener/gardener/pkg/apis/core/v1beta1/constants"
	gardenerhealthz "github.com/gardener/gardener/pkg/healthz"
	localinstall "github.com/gardener/gardener/pkg/provider-local/apis/local/install"
	localbackupbucket "github.com/gardener/gardener/pkg/provider-local/controller/backupbucket"
	localbackupentry "github.com/gardener/gardener/pkg/provider-local/controller/backupentry"
	"github.com/gardener/gardener/pkg/provider-local/controller/backupoptions"
	localcontrolplane "github.com/gardener/gardener/pkg/provider-local/controller/controlplane"
	localdnsrecord "github.com/gardener/gardener/pkg/provider-local/controller/dnsrecord"
	localhealthcheck "github.com/gardener/gardener/pkg/provider-local/controller/healthcheck"
	localinfrastructure "github.com/gardener/gardener/pkg/provider-local/controller/infrastructure"
	localingress "github.com/gardener/gardener/pkg/provider-local/controller/ingress"
	localservice "github.com/gardener/gardener/pkg/provider-local/controller/service"
	localworker "github.com/gardener/gardener/pkg/provider-local/controller/worker"
	"github.com/gardener/gardener/pkg/provider-local/local"
	controlplanewebhook "github.com/gardener/gardener/pkg/provider-local/webhook/controlplane"
	"github.com/gardener/gardener/pkg/utils/retry"
)

var hostIP string

func init() {
	addrs, err := net.InterfaceAddrs()
	utilruntime.Must(err)

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				hostIP = ipnet.IP.String()
				break
			}
		}
	}
}

// NewControllerManagerCommand creates a new command for running a local provider controller.
func NewControllerManagerCommand(ctx context.Context) *cobra.Command {
	var (
		restOpts = &extensionscmdcontroller.RESTOptions{}
		mgrOpts  = &extensionscmdcontroller.ManagerOptions{
			LeaderElection:             true,
			LeaderElectionResourceLock: resourcelock.LeasesResourceLock,
			LeaderElectionID:           extensionscmdcontroller.LeaderElectionNameID(local.Name),
			LeaderElectionNamespace:    os.Getenv("LEADER_ELECTION_NAMESPACE"),
			WebhookServerPort:          443,
			WebhookCertDir:             "/tmp/gardener-extensions-cert",
			MetricsBindAddress:         ":8080",
			HealthBindAddress:          ":8081",
		}
		generalOpts = &extensionscmdcontroller.GeneralOptions{}

		// options for the health care controller
		healthCheckCtrlOpts = &extensionscmdcontroller.ControllerOptions{
			MaxConcurrentReconciles: 5,
		}

		// options for the controlplane controller
		controlPlaneCtrlOpts = &extensionscmdcontroller.ControllerOptions{
			MaxConcurrentReconciles: 5,
		}

		// options for the dnsrecord controller
		dnsRecordCtrlOpts = &localdnsrecord.ControllerOptions{
			MaxConcurrentReconciles: 1,
		}

		// options for the ingress controller
		ingressCtrlOpts = &localingress.ControllerOptions{
			MaxConcurrentReconciles: 5,
		}

		// options for the service controller
		serviceCtrlOpts = &localservice.ControllerOptions{
			MaxConcurrentReconciles: 5,
			HostIP:                  hostIP,
			APIServerSNIEnabled:     true,
		}

		// options for the local backupbucket controller
		localBackupBucketOptions = &backupoptions.ControllerOptions{
			BackupBucketPath:   backupoptions.DefaultBackupPath,
			ContainerMountPath: backupoptions.DefaultContainerMountPath,
		}

		// options for the operatingsystemconfig controller
		operatingSystemConfigCtrlOpts = &extensionscmdcontroller.ControllerOptions{
			MaxConcurrentReconciles: 5,
		}

		// options for the infrastructure controller
		infraCtrlOpts = &extensionscmdcontroller.ControllerOptions{
			MaxConcurrentReconciles: 5,
		}
		reconcileOpts = &extensionscmdcontroller.ReconcilerOptions{}

		// options for the worker controller
		workerCtrlOpts = &extensionscmdcontroller.ControllerOptions{
			MaxConcurrentReconciles: 5,
		}
		workerReconcileOpts = &worker.Options{
			DeployCRDs: true,
		}
		workerCtrlOptsUnprefixed = extensionscmdcontroller.NewOptionAggregator(workerCtrlOpts, workerReconcileOpts)

		heartbeatCtrlOptions = &extensionsheartbeatcmd.Options{
			ExtensionName:        local.Name,
			RenewIntervalSeconds: 30,
			Namespace:            os.Getenv("LEADER_ELECTION_NAMESPACE"),
		}

		// options for the webhook server
		webhookServerOptions = &extensionscmdwebhook.ServerOptions{
			Namespace: os.Getenv("WEBHOOK_CONFIG_NAMESPACE"),
		}

		controllerSwitches = ControllerSwitchOptions()
		webhookSwitches    = WebhookSwitchOptions()
		webhookOptions     = extensionscmdwebhook.NewAddToManagerOptions(
			local.Name,
			genericactuator.ShootWebhooksResourceName,
			genericactuator.ShootWebhookNamespaceSelector(local.Type),
			webhookServerOptions,
			webhookSwitches,
		)

		aggOption = extensionscmdcontroller.NewOptionAggregator(
			restOpts,
			mgrOpts,
			generalOpts,
			extensionscmdcontroller.PrefixOption("controlplane-", controlPlaneCtrlOpts),
			extensionscmdcontroller.PrefixOption("dnsrecord-", dnsRecordCtrlOpts),
			extensionscmdcontroller.PrefixOption("infrastructure-", infraCtrlOpts),
			extensionscmdcontroller.PrefixOption("worker-", &workerCtrlOptsUnprefixed),
			extensionscmdcontroller.PrefixOption("ingress-", ingressCtrlOpts),
			extensionscmdcontroller.PrefixOption("service-", serviceCtrlOpts),
			extensionscmdcontroller.PrefixOption("backupbucket-", localBackupBucketOptions),
			extensionscmdcontroller.PrefixOption("operatingsystemconfig-", operatingSystemConfigCtrlOpts),
			extensionscmdcontroller.PrefixOption("healthcheck-", healthCheckCtrlOpts),
			extensionscmdcontroller.PrefixOption("heartbeat-", heartbeatCtrlOptions),
			controllerSwitches,
			reconcileOpts,
			webhookOptions,
		)
	)

	cmd := &cobra.Command{
		Use: fmt.Sprintf("%s-controller-manager", local.Name),

		RunE: func(cmd *cobra.Command, args []string) error {
			if err := aggOption.Complete(); err != nil {
				return fmt.Errorf("error completing options: %w", err)
			}

			if err := heartbeatCtrlOptions.Validate(); err != nil {
				return err
			}

			if workerReconcileOpts.Completed().DeployCRDs {
				if err := worker.ApplyMachineResourcesForConfig(ctx, restOpts.Completed().Config); err != nil {
					return fmt.Errorf("error ensuring the machine CRDs: %w", err)
				}
			}

			mgr, err := manager.New(restOpts.Completed().Config, mgrOpts.Completed().Options())
			if err != nil {
				return fmt.Errorf("could not instantiate manager: %w", err)
			}

			scheme := mgr.GetScheme()
			if err := controller.AddToScheme(scheme); err != nil {
				return fmt.Errorf("could not update manager scheme: %w", err)
			}
			if err := localinstall.AddToScheme(scheme); err != nil {
				return fmt.Errorf("could not update manager scheme: %w", err)
			}
			if err := vpaautoscalingv1.AddToScheme(scheme); err != nil {
				return fmt.Errorf("could not update manager scheme: %w", err)
			}
			if err := machinev1alpha1.AddToScheme(scheme); err != nil {
				return fmt.Errorf("could not update manager scheme: %w", err)
			}
			if err := druidv1alpha1.AddToScheme(scheme); err != nil {
				return fmt.Errorf("could not update manager scheme: %w", err)
			}
			// add common meta types to schema for controller-runtime to use v1.ListOptions
			metav1.AddToGroupVersion(scheme, machinev1alpha1.SchemeGroupVersion)

			controlPlaneCtrlOpts.Completed().Apply(&localcontrolplane.DefaultAddOptions.Controller)
			dnsRecordCtrlOpts.Completed().Apply(&localdnsrecord.DefaultAddOptions)
			healthCheckCtrlOpts.Completed().Apply(&localhealthcheck.DefaultAddOptions.Controller)
			infraCtrlOpts.Completed().Apply(&localinfrastructure.DefaultAddOptions.Controller)
			operatingSystemConfigCtrlOpts.Completed().Apply(&oscommon.DefaultAddOptions.Controller)
			ingressCtrlOpts.Completed().Apply(&localingress.DefaultAddOptions)
			serviceCtrlOpts.Completed().Apply(&localservice.DefaultAddOptions)
			workerCtrlOpts.Completed().Apply(&localworker.DefaultAddOptions.Controller)
			localBackupBucketOptions.Completed().Apply(&localbackupbucket.DefaultAddOptions)
			localBackupBucketOptions.Completed().Apply(&localbackupentry.DefaultAddOptions)
			heartbeatCtrlOptions.Completed().Apply(&heartbeat.DefaultAddOptions)

			reconcileOpts.Completed().Apply(&localcontrolplane.DefaultAddOptions.IgnoreOperationAnnotation)
			reconcileOpts.Completed().Apply(&localdnsrecord.DefaultAddOptions.IgnoreOperationAnnotation)
			reconcileOpts.Completed().Apply(&localinfrastructure.DefaultAddOptions.IgnoreOperationAnnotation)
			reconcileOpts.Completed().Apply(&oscommon.DefaultAddOptions.IgnoreOperationAnnotation)
			reconcileOpts.Completed().Apply(&localworker.DefaultAddOptions.IgnoreOperationAnnotation)

			// TODO(rfranzke): Remove this in a future version.
			localworker.DefaultAddOptions.ManageMCM = !generalOpts.Completed().GardenletManagesMCM
			controlplanewebhook.ManageMCM = !generalOpts.Completed().GardenletManagesMCM

			if err := mgr.AddReadyzCheck("informer-sync", gardenerhealthz.NewCacheSyncHealthz(mgr.GetCache())); err != nil {
				return fmt.Errorf("could not add readycheck for informers: %w", err)
			}
			if err := mgr.AddHealthzCheck("ping", healthz.Ping); err != nil {
				return fmt.Errorf("could not add healthcheck: %w", err)
			}
			if err := mgr.AddReadyzCheck("webhook-server", mgr.GetWebhookServer().StartedChecker()); err != nil {
				return fmt.Errorf("could not add readycheck of webhook to manager: %w", err)
			}

			atomicShootWebhookConfig, err := webhookOptions.Completed().AddToManager(ctx, mgr)
			if err != nil {
				return fmt.Errorf("could not add webhooks to manager: %w", err)
			}
			localcontrolplane.DefaultAddOptions.ShootWebhookConfig = atomicShootWebhookConfig
			localcontrolplane.DefaultAddOptions.WebhookServerNamespace = webhookOptions.Server.Namespace

			// Send empty patches on start-up to trigger webhooks
			if err := mgr.Add(&webhookTriggerer{client: mgr.GetClient()}); err != nil {
				return fmt.Errorf("error adding runnable for triggering DNS config webhook: %w", err)
			}

			// TODO(rfranzke): Remove this block after v1.71 got released.
			// Migrate existing machine pods to new NetworkPolicy labels to make upgrade e2e tests work.
			{
				if err := mgr.Add(manager.RunnableFunc(func(ctx context.Context) error {
					machinePods := &corev1.PodList{}
					if err := mgr.GetClient().List(ctx, machinePods, client.MatchingLabels{
						"app":              "machine",
						"machine-provider": "local",
					}); err != nil {
						return err
					}

					for _, p := range machinePods.Items {
						pod := p
						patch := client.MergeFrom(pod.DeepCopy())
						metav1.SetMetaDataLabel(&pod.ObjectMeta, "networking.gardener.cloud/to-runtime-apiserver", "allowed")
						metav1.SetMetaDataLabel(&pod.ObjectMeta, "networking.resources.gardener.cloud/to-kube-apiserver-tcp-443", "allowed")
						if err := mgr.GetClient().Patch(ctx, &pod, patch); err != nil {
							return err
						}
					}

					shootNamespaces := &corev1.NamespaceList{}
					if err := mgr.GetClient().List(ctx, shootNamespaces, client.MatchingLabels{v1beta1constants.GardenRole: v1beta1constants.GardenRoleShoot}); err != nil {
						return err
					}

					for _, namespace := range shootNamespaces.Items {
						service := &corev1.Service{
							ObjectMeta: metav1.ObjectMeta{
								Name:      "machines",
								Namespace: namespace.Name,
							},
							Spec: corev1.ServiceSpec{
								Type:      corev1.ServiceTypeClusterIP,
								ClusterIP: corev1.ClusterIPNone,
								Selector: map[string]string{
									"app":              "machine",
									"machine-provider": "local",
								},
								Ports: []corev1.ServicePort{{
									Port:       10250,
									Protocol:   corev1.ProtocolTCP,
									TargetPort: intstr.FromInt(10250),
								}},
							},
						}

						if err := mgr.GetClient().Create(ctx, service); client.IgnoreAlreadyExists(err) != nil {
							return err
						}
					}

					return nil
				})); err != nil {
					return fmt.Errorf("error adding runnable for machine pod network policy label migration: %w", err)
				}
			}

			if err := controllerSwitches.Completed().AddToManager(mgr); err != nil {
				return fmt.Errorf("could not add controllers to manager: %w", err)
			}

			mgr.GetLogger().Info("Started with", "hostIP", serviceCtrlOpts.HostIP)

			if err := mgr.Start(ctx); err != nil {
				return fmt.Errorf("error running manager: %w", err)
			}

			return nil
		},
	}

	aggOption.AddFlags(cmd.Flags())

	return cmd
}

type webhookTriggerer struct {
	client client.Client
}

func (w *webhookTriggerer) NeedLeaderElection() bool {
	return true
}

func (w *webhookTriggerer) Start(ctx context.Context) error {
	// Wait for the reconciler to populate the webhook CA into the configurations before triggering the webhooks.
	timeoutCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	if err := retry.Until(timeoutCtx, time.Second, func(ctx context.Context) (bool, error) {
		webhookConfig := &admissionregistrationv1.MutatingWebhookConfiguration{ObjectMeta: metav1.ObjectMeta{Name: "gardener-extension-" + local.Name}}
		if err := w.client.Get(ctx, client.ObjectKeyFromObject(webhookConfig), webhookConfig); err != nil {
			if !apierrors.IsNotFound(err) {
				return retry.SevereError(err)
			}
			return retry.MinorError(fmt.Errorf("webhook was not yet created"))
		}

		for _, webhook := range webhookConfig.Webhooks {
			// We can return when we find the first webhook w/o CA bundle since the reconciler would populate it into
			// all webhooks at the same time.
			if len(webhook.ClientConfig.CABundle) == 0 {
				return retry.MinorError(fmt.Errorf("CA bundle was not yet populated to all webhooks"))
			}
		}

		return retry.Ok()
	}); err != nil {
		return err
	}

	if err := w.trigger(ctx, w.client, nil, w.client.Status(), &corev1.NodeList{}); err != nil {
		return err
	}

	return w.trigger(ctx, w.client, w.client, nil, &appsv1.DeploymentList{}, client.MatchingLabels{"app": "dependency-watchdog-prober"})
}

func (w *webhookTriggerer) trigger(ctx context.Context, reader client.Reader, writer client.Writer, statusWriter client.StatusWriter, objectList client.ObjectList, opts ...client.ListOption) error {
	if err := reader.List(ctx, objectList, opts...); err != nil {
		return err
	}

	return meta.EachListItem(objectList, func(obj runtime.Object) error {
		switch object := obj.(type) {
		case *appsv1.Deployment:
			return writer.Patch(ctx, object, client.RawPatch(types.StrategicMergePatchType, []byte("{}")))
		case *corev1.Node:
			return statusWriter.Patch(ctx, object, client.RawPatch(types.StrategicMergePatchType, []byte("{}")))
		}
		return nil
	})
}
