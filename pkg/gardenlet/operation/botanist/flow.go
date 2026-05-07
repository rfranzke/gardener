// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package botanist

import (
	"context"

	v1beta1helper "github.com/gardener/gardener/pkg/api/core/v1beta1/helper"
	v1beta1constants "github.com/gardener/gardener/pkg/apis/core/v1beta1/constants"
	gardenerextensions "github.com/gardener/gardener/pkg/extensions"
	"github.com/gardener/gardener/pkg/utils/flow"
	gardenerutils "github.com/gardener/gardener/pkg/utils/gardener"
)

const (
	GroupDeployNamespaces flow.TaskID= "GroupDeployNamespaces"
	GroupDeployCredentials flow.TaskID = "GroupDeployCredentials"
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
		Name:         "Deploying cloud provider account secret",
		Fn:           b.DeployCloudProviderSecret,
		SkipIf:       b.Shoot.Credentials == nil,
	}).WithDependencies(GroupDeployNamespaces)
}

func (b *Botanist) ReconcileClusterResource(g *flow.Graph, dependencies ...flow.TaskIDer) flow.TaskIDs {
	return flow.NewTaskIDs(g.Add(flow.Task{
		Name: "Reconciling extensions.gardener.cloud/v1alpha1.Cluster resource",
		Fn: func(ctx context.Context) error {
			return gardenerextensions.SyncClusterResourceToSeed(ctx, b.SeedClientSet.Client(), b.Shoot.ControlPlaneNamespace, b.Shoot.GetInfo(), b.Shoot.CloudProfile, nil)
		},
		Dependencies: flow.NewTaskIDs(dependencies...),
	}))
}
