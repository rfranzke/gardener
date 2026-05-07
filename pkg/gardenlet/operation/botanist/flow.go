// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package botanist

import (
	"context"

	v1beta1helper "github.com/gardener/gardener/pkg/api/core/v1beta1/helper"
	v1beta1constants "github.com/gardener/gardener/pkg/apis/core/v1beta1/constants"
	seedsystem "github.com/gardener/gardener/pkg/component/seed/system"
	gardenerextensions "github.com/gardener/gardener/pkg/extensions"
	"github.com/gardener/gardener/pkg/utils/flow"
	gardenerutils "github.com/gardener/gardener/pkg/utils/gardener"
)

const (
	GroupDeployNamespaces         flow.TaskID = "GroupDeployNamespaces"
	GroupDeployCredentials        flow.TaskID = "GroupDeployCredentials"
	GroupReconcileClusterResource flow.TaskID = "GroupReconcileClusterResource"
	GroupInitializeSecretsManager flow.TaskID = "GroupInitializeSecretsManager"
)

func (b *Botanist) DeployNamespaces() flow.TaskGroup {
	return flow.NewTaskGroup(
		GroupDeployNamespaces,
		flow.Task{
			Name: "Deploying control plane namespace",
			Fn:   b.DeployControlPlaneNamespace,
		},
		flow.Task{
			Name: "Deploying garden namespace",
			Fn: func(ctx context.Context) error {
				return gardenerutils.ReconcileGardenNamespace(ctx, b.SeedClientSet.Client(), v1beta1constants.GardenNamespace, v1beta1helper.ControlPlaneWorkerPoolForShoot(b.Shoot.GetInfo().Spec.Provider.Workers).Zones, true, nil)
			},
			SkipIf: !b.Shoot.RunsControlPlane(),
		},
	)
}

func (b *Botanist) DeployCloudProviderCredentials() flow.TaskGroup {
	return flow.NewTaskGroup(GroupDeployCredentials, flow.Task{
		Name:   "Deploying cloud provider account secret",
		Fn:     b.DeployCloudProviderSecret,
		SkipIf: b.Shoot.Credentials == nil,
	}).WithDependencies(GroupDeployNamespaces)
}

func (b *Botanist) ReconcileClusterResource() flow.TaskGroup {
	return flow.NewTaskGroup(GroupReconcileClusterResource, flow.Task{
		Name: "Reconciling extensions.gardener.cloud/v1alpha1.Cluster resource",
		Fn: func(ctx context.Context) error {
			return gardenerextensions.SyncClusterResourceToSeed(ctx, b.SeedClientSet.Client(), b.Shoot.ControlPlaneNamespace, b.Shoot.GetInfo(), b.Shoot.CloudProfile, b.GetSeed())
		},
	}).WithDependencies(GroupReconcileCRDs)
}

func (b *Botanist) InitializeSecretsManager() flow.TaskGroup {
	return flow.NewTaskGroup(GroupInitializeSecretsManager, flow.Task{
		Name: "Initializing secrets management",
		Fn:   b.InitializeSecretsManagement,
	}).WithDependencies(GroupReconcileClusterResource)
}

func (b *Botanist) InitializeResourceManager(podNetworkAvailable, shootIsGarden bool) flow.TaskGroup {
	g := flow.NewTaskGroup(GroupReconcileClusterResource).
		WithDependencies(GroupDeployNamespaces, GroupInitializeSecretsManager, GroupReconcileCRDs)

	var (
		gardenadmBootstrap = b.Shoot.IsSelfHosted() && !b.Shoot.RunsControlPlane()

		deployGardenerResourceManager = g.Add(flow.Task{
			Name: "Deploying gardener-resource-manager",
			Fn: func(ctx context.Context) error {
				b.Shoot.Components.ControlPlane.RuntimeResourceManager.SetBootstrapControlPlaneNode(!podNetworkAvailable)
				b.Shoot.Components.ControlPlane.ResourceManager.SetBootstrapControlPlaneNode(!podNetworkAvailable)

				if shootIsGarden || (gardenadmBootstrap) {
					return b.Shoot.Components.ControlPlane.ResourceManager.Deploy(ctx)
				}

				return flow.Parallel(
					b.Shoot.Components.ControlPlane.RuntimeResourceManager.Deploy,
					b.Shoot.Components.ControlPlane.ResourceManager.Deploy,
				)(ctx)
			},
		})
		waitUntilGardenerResourceManagerReady = g.Add(flow.Task{
			Name: "Waiting until gardener-resource-manager reports readiness",
			Fn: func(ctx context.Context) error {
				if shootIsGarden || (gardenadmBootstrap) {
					return b.Shoot.Components.ControlPlane.ResourceManager.Wait(ctx)
				}

				return flow.Parallel(
					b.Shoot.Components.ControlPlane.RuntimeResourceManager.Wait,
					b.Shoot.Components.ControlPlane.ResourceManager.Wait,
				)(ctx)
			},
			Dependencies: flow.NewTaskIDs(deployGardenerResourceManager),
		})
		_ = g.Add(flow.Task{
			Name: "Deploying seed system resources",
			Fn: func(ctx context.Context) error {
				return seedsystem.New(b.SeedClientSet.Client(), b.Shoot.ControlPlaneNamespace, seedsystem.Values{}).Deploy(ctx)
			},
			Dependencies: flow.NewTaskIDs(waitUntilGardenerResourceManagerReady),
		})
		_ = g.Add(flow.Task{
			Name:         "Deploying shoot system resources",
			Fn:           b.DeployShootSystem,
			SkipIf:       gardenadmBootstrap,
			Dependencies: flow.NewTaskIDs(waitUntilGardenerResourceManagerReady),
		})
	)

	return g
}
