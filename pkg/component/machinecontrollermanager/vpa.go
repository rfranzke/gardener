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

package machinecontrollermanager

import (
	v1 "k8s.io/api/core/v1"
	vpaautoscalingv1 "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/apis/autoscaling.k8s.io/v1"
)

// GetSidecarContainer returns a corev1.Container object which is
// injected into the mcm deployment managed by the gardenlet. This
// function can be used in provider specific
// machine-controller-manager implementations, when a meaningful
// default sidecar container is required. Thus, this function returns
// kind of a template for these scenarios.
func GetVPAContainerPolicy(namespace, providerName string, minAllowed, maxAllowed v1.ResourceList) vpaautoscalingv1.ContainerResourcePolicy {

	var ccv = vpaautoscalingv1.ContainerControlledValuesRequestsOnly
	return vpaautoscalingv1.ContainerResourcePolicy{
		ContainerName:    "machine-controller-manager-" + providerName,
		ControlledValues: &ccv,
		MinAllowed:       minAllowed,
		MaxAllowed:       maxAllowed,
	}
}
