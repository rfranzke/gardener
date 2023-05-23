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

package controlplane

import (
	"bytes"
	"context"
	_ "embed"
	"text/template"

	"github.com/Masterminds/semver"
	"github.com/Masterminds/sprig"
	"github.com/go-logr/logr"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	kubeletconfigv1beta1 "k8s.io/kubelet/config/v1beta1"
	"k8s.io/utils/pointer"

	"github.com/gardener/gardener/extensions/pkg/webhook"
	extensionscontextwebhook "github.com/gardener/gardener/extensions/pkg/webhook/context"
	"github.com/gardener/gardener/extensions/pkg/webhook/controlplane/genericmutator"
	v1beta1constants "github.com/gardener/gardener/pkg/apis/core/v1beta1/constants"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	"github.com/gardener/gardener/pkg/component/machinecontrollermanager"
	"github.com/gardener/gardener/pkg/provider-local/imagevector"
	"github.com/gardener/gardener/pkg/utils"
	gardenerutils "github.com/gardener/gardener/pkg/utils/gardener"
	vpaautoscalingv1 "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/apis/autoscaling.k8s.io/v1"
)

const pathContainerdConfigScript = v1beta1constants.OperatingSystemConfigFilePathBinaries + "/init-containerd-with-registry-mirrors"

var (
	tplNameInitializer = "init"
	//go:embed templates/scripts/configure-containerd.tpl.sh
	tplContentInitializer string
	tplInitializer        *template.Template
)

func init() {
	var err error
	tplInitializer, err = template.
		New(tplNameInitializer).
		Funcs(sprig.TxtFuncMap()).
		Parse(tplContentInitializer)
	utilruntime.Must(err)
}

// NewEnsurer creates a new controlplane ensurer.
func NewEnsurer(logger logr.Logger, manageMCM bool) genericmutator.Ensurer {
	return &ensurer{
		logger:    logger.WithName("local-controlplane-ensurer"),
		manageMCM: manageMCM,
	}
}

type ensurer struct {
	genericmutator.NoopEnsurer
	logger    logr.Logger
	manageMCM bool
}

// EnsureMachineControllerManagerDeployment ensures that the machine-controller-manager deployment conforms to the provider requirements.
func (e *ensurer) EnsureMachineControllerManagerDeployment(ctx context.Context, gctx extensionscontextwebhook.GardenContext, newObj, _ *appsv1.Deployment) error {
	if e.manageMCM {
		return nil
	}

	image, err := imagevector.ImageVector().FindImage(imagevector.ImageNameMachineControllerManagerProviderLocal)
	if err != nil {
		return err
	}

	container := machinecontrollermanager.GetSidecarContainer(newObj.Namespace, "provider-local", image.String())
	newObj.Spec.Template.Spec.Containers = webhook.EnsureContainerWithName(newObj.Spec.Template.Spec.Containers, container)

	return nil
}

// EnsureMachineControllerManagerVPA ensures that the machine-controller-manager VPA conforms to the provider requirements.
func (e *ensurer) EnsureMachineControllerManagerVPA(ctx context.Context, gctx extensionscontextwebhook.GardenContext, newObj, _ *vpaautoscalingv1.VerticalPodAutoscaler) error {
	if e.manageMCM {
		return nil
	}

	minAllowed := v1.ResourceList{
		v1.ResourceMemory: resource.MustParse("64Mi"),
	}
	maxAllowed := v1.ResourceList{
		v1.ResourceCPU:    resource.MustParse("2"),
		v1.ResourceMemory: resource.MustParse("5G"),
	}

	containerResourcePolicy := machinecontrollermanager.GetVPAContainerPolicy(newObj.Namespace, "provider-local", minAllowed, maxAllowed)
	newObj.Spec.ResourcePolicy.ContainerPolicies = webhook.EnsureContainerResourcePolicyWithName(newObj.Spec.ResourcePolicy.ContainerPolicies, containerResourcePolicy)

	return nil
}

func (e *ensurer) EnsureKubeAPIServerDeployment(_ context.Context, _ extensionscontextwebhook.GardenContext, new, _ *appsv1.Deployment) error {
	metav1.SetMetaDataLabel(&new.Spec.Template.ObjectMeta, gardenerutils.NetworkPolicyLabel("machines", 10250), v1beta1constants.LabelNetworkPolicyAllowed)
	return nil
}

func (e *ensurer) EnsureKubeletConfiguration(_ context.Context, _ extensionscontextwebhook.GardenContext, _ *semver.Version, newObj, _ *kubeletconfigv1beta1.KubeletConfiguration) error {
	newObj.FailSwapOn = pointer.Bool(false)
	newObj.CgroupDriver = "systemd"
	return nil
}

func (e *ensurer) EnsureAdditionalFiles(ctx context.Context, gc extensionscontextwebhook.GardenContext, new, _ *[]extensionsv1alpha1.File) error {
	var script bytes.Buffer
	if err := tplInitializer.Execute(&script, nil); err != nil {
		return err
	}

	appendUniqueFile(new, extensionsv1alpha1.File{
		Path:        pathContainerdConfigScript,
		Permissions: pointer.Int32(0744),
		Content: extensionsv1alpha1.FileContent{
			Inline: &extensionsv1alpha1.FileContentInline{
				Encoding: "b64",
				Data:     utils.EncodeBase64(script.Bytes()),
			},
		},
	})
	return nil
}

func (e *ensurer) EnsureAdditionalUnits(_ context.Context, _ extensionscontextwebhook.GardenContext, new, _ *[]extensionsv1alpha1.Unit) error {
	const unitNameInitializer = "containerd-configuration-local-setup.service"
	unit := extensionsv1alpha1.Unit{
		Name:    unitNameInitializer,
		Command: pointer.String("start"),
		Enable:  pointer.Bool(true),
		Content: pointer.String(`[Unit]
Description=Containerd config configuration for local-setup

[Install]
WantedBy=multi-user.target

[Unit]
After=containerd-initializer.service
Requires=containerd-initializer.service

[Service]
Type=oneshot
RemainAfterExit=yes
ExecStart=` + pathContainerdConfigScript)}

	appendUniqueUnit(new, unit)

	return nil
}

// appendUniqueFile appends a unit file only if it does not exist, otherwise overwrite content of previous files
func appendUniqueFile(files *[]extensionsv1alpha1.File, file extensionsv1alpha1.File) {
	resFiles := make([]extensionsv1alpha1.File, 0, len(*files))

	for _, f := range *files {
		if f.Path != file.Path {
			resFiles = append(resFiles, f)
		}
	}

	*files = append(resFiles, file)
}

// appendUniqueUnit appends a unit only if it does not exist, otherwise overwrite content of previous unit
func appendUniqueUnit(units *[]extensionsv1alpha1.Unit, unit extensionsv1alpha1.Unit) {
	resFiles := make([]extensionsv1alpha1.Unit, 0, len(*units))

	for _, f := range *units {
		if f.Name != unit.Name {
			resFiles = append(resFiles, f)
		}
	}

	*units = append(resFiles, unit)
}
