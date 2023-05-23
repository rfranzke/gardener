// Copyright 2023 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
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

package botanist

import (
	"context"

	machinev1alpha1 "github.com/gardener/machine-controller-manager/pkg/apis/machine/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	gardencorev1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	v1beta1constants "github.com/gardener/gardener/pkg/apis/core/v1beta1/constants"
	"github.com/gardener/gardener/pkg/client/kubernetes"
	"github.com/gardener/gardener/pkg/component/machinecontrollermanager"
	"github.com/gardener/gardener/pkg/utils/images"
	"github.com/gardener/gardener/pkg/utils/imagevector"
	kubernetesutils "github.com/gardener/gardener/pkg/utils/kubernetes"
)

// DefaultMachineControllerManager returns a deployer for the machine-controller-manager.
func (b *Botanist) DefaultMachineControllerManager(ctx context.Context) (machinecontrollermanager.Interface, error) {
	image, err := b.ImageVector.FindImage(images.ImageNameMachineControllerManager, imagevector.RuntimeVersion(b.SeedVersion()), imagevector.TargetVersion(b.ShootVersion()))
	if err != nil {
		return nil, err
	}

	machineDeploymentList := &machinev1alpha1.MachineDeploymentList{}
	if err := b.SeedClientSet.Client().List(ctx, machineDeploymentList, client.InNamespace(b.Shoot.SeedNamespace)); err != nil {
		return nil, err
	}

	var replicas int32 = 1
	switch {
	// if there are any existing machine deployments present with a positive replica count then MCM is needed.
	case machineDeploymentWithPositiveReplicaCountExist(machineDeploymentList):
		replicas = 1
	// If the cluster is hibernated then there is no further need of MCM and therefore its desired replicas is 0
	case b.Shoot.HibernationEnabled && b.Shoot.GetInfo().Status.IsHibernated:
		replicas = 0
	// If the cluster is created with hibernation enabled, then desired replicas for MCM is 0
	case b.Shoot.HibernationEnabled && (b.Shoot.GetInfo().Status.LastOperation == nil || b.Shoot.GetInfo().Status.LastOperation.Type == gardencorev1beta1.LastOperationTypeCreate):
		replicas = 0
	// If shoot is either waking up or in the process of hibernation then, MCM is required and therefore its desired replicas is 1
	case b.Shoot.HibernationEnabled != b.Shoot.GetInfo().Status.IsHibernated:
		replicas = 1
	}

	return machinecontrollermanager.New(
		b.SeedClientSet.Client(),
		b.Shoot.SeedNamespace,
		b.SecretsManager,
		machinecontrollermanager.Values{
			Image:                    image.String(),
			Replicas:                 replicas,
			RuntimeKubernetesVersion: b.Seed.KubernetesVersion,
		},
	), nil
}

// ScaleMachineControllerManagerToZero scales machine-controller-manager replicas to zero.
func (b *Botanist) ScaleMachineControllerManagerToZero(ctx context.Context) error {
	return kubernetes.ScaleDeployment(ctx, b.SeedClientSet.Client(), kubernetesutils.Key(b.Shoot.SeedNamespace, v1beta1constants.DeploymentNameMachineControllerManager), 0)
}

func machineDeploymentWithPositiveReplicaCountExist(existingMachineDeployments *machinev1alpha1.MachineDeploymentList) bool {
	for _, machineDeployment := range existingMachineDeployments.Items {
		if machineDeployment.Status.Replicas > 0 {
			return true
		}
	}
	return false
}
