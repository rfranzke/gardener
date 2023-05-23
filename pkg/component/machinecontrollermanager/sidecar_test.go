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
	corev1 "k8s.io/api/core/v1"

	. "github.com/gardener/gardener/pkg/component/machinecontrollermanager"
	"github.com/gardener/gardener/pkg/utils"
	gardenerutils "github.com/gardener/gardener/pkg/utils/gardener"
)

var _ = Describe("Sidecar", func() {

	It("Should return a default provider specific mcm container spec", func() {

		namespace := "test-namespace"
		provider := "provider-test"
		image := "provider-test:latest"

		container := GetSidecarContainer(namespace, provider, image)
		Expect(container).To(Equal(
			corev1.Container{
				Name:                     "machine-controller-manager-" + provider,
				Image:                    image,
				TerminationMessagePath:   "/dev/termination-log",
				TerminationMessagePolicy: "File",
				ImagePullPolicy:          "IfNotPresent",
				Command: []string{
					"./machine-controller",
					"--control-kubeconfig=inClusterConfig",
					"--machine-creation-timeout=20m",
					"--machine-drain-timeout=2h",
					"--machine-health-timeout=10m",
					"--machine-safety-apiserver-statuscheck-timeout=30s",
					"--machine-safety-apiserver-statuscheck-period=1m",
					"--machine-safety-orphan-vms-period=30m",
					"--namespace=" + namespace,
					"--port=10259",
					"--target-kubeconfig=" + gardenerutils.PathGenericKubeconfig,
					"--v=3",
				},
				VolumeMounts: []corev1.VolumeMount{
					{
						Name:      "kubeconfig",
						MountPath: gardenerutils.VolumeMountPathGenericKubeconfig,
						ReadOnly:  true,
					},
				},
				LivenessProbe: &corev1.Probe{
					ProbeHandler: corev1.ProbeHandler{
						HTTPGet: &corev1.HTTPGetAction{
							Path:   "/healthz",
							Port:   *utils.IntStrPtrFromString("10259"),
							Scheme: "http",
						},
					},
					InitialDelaySeconds: 30,
					TimeoutSeconds:      5,
					PeriodSeconds:       10,
					SuccessThreshold:    1,
					FailureThreshold:    3,
				},
			},
		))
	})
})
