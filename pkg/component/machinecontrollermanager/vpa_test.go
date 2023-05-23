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

package machinecontrollermanager_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/gardener/gardener/pkg/component/machinecontrollermanager"
	corev1 "k8s.io/api/core/v1"

	"k8s.io/apimachinery/pkg/api/resource"
	vpaautoscalingv1 "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/apis/autoscaling.k8s.io/v1"
)

var _ = Describe("Vpa", func() {

	It("Should return a valid VPA container policy", func() {

		namespace := "test-namespace"
		provider := "provider-test"

		minAllowed := corev1.ResourceList{corev1.ResourceMemory: resource.MustParse("1M")}
		maxAllowed := corev1.ResourceList{corev1.ResourceMemory: resource.MustParse("5M")}

		containerResourcePolicy := GetVPAContainerPolicy(namespace, provider, minAllowed, maxAllowed)

		var ccv = vpaautoscalingv1.ContainerControlledValuesRequestsOnly
		Expect(containerResourcePolicy).To(Equal(
			vpaautoscalingv1.ContainerResourcePolicy{
				ContainerName:    "machine-controller-manager-" + provider,
				ControlledValues: &ccv,
				MinAllowed:       minAllowed,
				MaxAllowed:       maxAllowed,
			},
		))
	})
})
