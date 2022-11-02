// Copyright (c) 2022 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	componentbaseconfigv1alpha1 "k8s.io/component-base/config/v1alpha1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ResourceManagerConfiguration defines the configuration for the gardener-resource-manager.
type ResourceManagerConfiguration struct {
	metav1.TypeMeta `json:",inline"`
	// SourceClientConnection specifies the client connection settings for the proxy server
	// to use when communicating with the source apiserver.
	// +optional
	SourceClientConnection SourceClientConnection `json:"sourceClientConnection"`
	// TargetClientConnection specifies the client connection settings for the proxy server
	// to use when communicating with the target apiserver.
	// +optional
	TargetClientConnection *TargetClientConnection `json:"targetClientConnection,omitempty"`
	// LeaderElection defines the configuration of leader election client.
	LeaderElection componentbaseconfigv1alpha1.LeaderElectionConfiguration `json:"leaderElection"`
	// Server defines the configuration of the HTTP server.
	Server ServerConfiguration `json:"server"`
	// Debugging holds configuration for Debugging related features.
	// +optional
	Debugging *componentbaseconfigv1alpha1.DebuggingConfiguration `json:"debugging,omitempty"`
	// LogLevel is the level/severity for the logs. Must be one of [info,debug,error].
	LogLevel string `json:"logLevel"`
	// LogFormat is the output format for the logs. Must be one of [text,json].
	LogFormat string `json:"logFormat"`
	// Controllers defines the configuration of the controllers.
	Controllers ResourceManagerControllerConfiguration `json:"controllers"`
	// Webhooks defines the configuration of the webhooks.
	Webhooks ResourceManagerWebhookConfiguration `json:"webhooks"`
}

// SourceClientConnection specifies the client connection settings
// for the proxy server to use when communicating with the seed apiserver.
type SourceClientConnection struct {
	componentbaseconfigv1alpha1.ClientConnectionConfiguration `json:",inline"`
	// Namespace in which the ManagedResources should be observed (defaults to "all namespaces").
	// +optional
	Namespace *string `json:"namespace,omitempty"`
	// CacheResyncPeriod specifies the duration how often the cache for the source cluster is resynced.
	// +optional
	CacheResyncPeriod *metav1.Duration `json:"cacheResyncPeriod,omitempty"`
}

// TargetClientConnection specifies the client connection settings
// for the proxy server to use when communicating with the shoot apiserver.
type TargetClientConnection struct {
	componentbaseconfigv1alpha1.ClientConnectionConfiguration `json:",inline"`
	// Namespace in which controllers for the target clusters act on objects (defaults to "all namespaces").
	// +optional
	Namespace *string `json:"namespace,omitempty"`
	// DisableCachedClient specifies whether the cache for the target cluster client should be disabled. If true, then
	// each request is performed with a direct client.
	// +optional
	DisableCachedClient *bool `json:"disableCachedClient,omitempty"`
	// CacheResyncPeriod specifies the duration how often the cache for the target cluster is resynced.
	// +optional
	CacheResyncPeriod *metav1.Duration `json:"cacheResyncPeriod,omitempty"`
}

// ServerConfiguration contains details for the HTTP(S) servers.
type ServerConfiguration struct {
	// Webhooks is the configuration for the HTTPS webhook server.
	Webhooks HTTPSServer `json:"webhooks"`
	// HealthProbes is the configuration for serving the healthz and readyz endpoints.
	// +optional
	HealthProbes *Server `json:"healthProbes,omitempty"`
	// Metrics is the configuration for serving the metrics endpoint.
	// +optional
	Metrics *Server `json:"metrics,omitempty"`
}

// Server contains information for HTTP(S) server configuration.
type Server struct {
	// BindAddress is the IP address on which to listen for the specified port.
	BindAddress string `json:"bindAddress"`
	// Port is the port on which to serve unsecured, unauthenticated access.
	Port int `json:"port"`
}

// HTTPSServer is the configuration for the HTTPSServer server.
type HTTPSServer struct {
	// Server is the configuration for the bind address and the port.
	Server `json:",inline"`
	// TLSServer contains information about the TLS configuration for an HTTPS server.
	TLS TLSServer `json:"tls"`
}

// TLSServer contains information about the TLS configuration for an HTTPS server.
type TLSServer struct {
	// ServerCertDir is the path to a directory containing the server's TLS certificate and key (the files must be
	// named tls.crt and tls.key respectively).
	ServerCertDir string `json:"serverCertDir"`
}

// ResourceManagerControllerConfiguration defines the configuration of the controllers.
type ResourceManagerControllerConfiguration struct {
	// ClusterID is the ID of the source cluster.
	// +optional
	ClusterID *string `json:"clusterID,omitempty"`
	// ResourceClass is the name of the class in ManagedResources to filter for.
	// +optional
	ResourceClass *string `json:"resourceClass,omitempty"`

	// KubeletCSRApprover is the configuration for the kubelet-csr-approver controller.
	KubeletCSRApprover KubeletCSRApproverControllerConfig `json:"kubeletCSRApprover"`
	// GarbageCollector is the configuration for the garbage-collector controller.
	GarbageCollector GarbageCollectorControllerConfig `json:"garbageCollector"`
	// Health is the configuration for the health controller.
	Health HealthControllerConfig `json:"health"`
	// ManagedResource is the configuration for the managed resource controller.
	ManagedResource ManagedResourceControllerConfig `json:"managedResource"`
	// RootCAPublisher is the configuration for the root-ca-publisher controller.
	RootCAPublisher RootCAPublisherControllerConfig `json:"rootCAPublisher"`
	// Secret is the configuration for the secret controller.
	Secret SecretControllerConfig `json:"secret"`
	// TokenInvalidator is the configuration for the token-invalidator controller.
	TokenInvalidator TokenInvalidatorControllerConfig `json:"tokenInvalidator"`
	// TokenRequestor is the configuration for the token-requestor controller.
	TokenRequestor TokenRequestorControllerConfig `json:"tokenRequestor"`
}

// KubeletCSRApproverControllerConfig is the configuration for the kubelet-csr-approver controller.
type KubeletCSRApproverControllerConfig struct {
	// Enabled defines whether this controller is enabled.
	Enabled bool `json:"enabled"`
	// ConcurrentSyncs is the number of concurrent worker routines for this controller.
	// +optional
	ConcurrentSyncs *int `json:"concurrentSyncs,omitempty"`
}

// GarbageCollectorControllerConfig is the configuration for the garbage-collector controller.
type GarbageCollectorControllerConfig struct {
	// Enabled defines whether this controller is enabled.
	Enabled bool `json:"enabled"`
	// SyncPeriod is the duration how often the controller performs its reconciliation.
	// +optional
	SyncPeriod *metav1.Duration `json:"syncPeriod,omitempty"`
}

// HealthControllerConfig is the configuration for the health controller.
type HealthControllerConfig struct {
	// ConcurrentSyncs is the number of concurrent worker routines for this controller.
	// +optional
	ConcurrentSyncs *int `json:"concurrentSyncs,omitempty"`
	// SyncPeriod is the duration how often the controller performs its reconciliation.
	// +optional
	SyncPeriod *metav1.Duration `json:"syncPeriod,omitempty"`
}

// ManagedResourceControllerConfig is the configuration for the managed resource controller.
type ManagedResourceControllerConfig struct {
	// ConcurrentSyncs is the number of concurrent worker routines for this controller.
	// +optional
	ConcurrentSyncs *int `json:"concurrentSyncs,omitempty"`
	// SyncPeriod is the duration how often the controller performs its reconciliation.
	// +optional
	SyncPeriod *metav1.Duration `json:"syncPeriod,omitempty"`
	// AlwaysUpdate specifies whether resources will only be updated if their desired state differs from the actual
	// state. If true, an update request will be sent in each reconciliation independent of this condition.
	// +optional
	AlwaysUpdate *bool `json:"alwaysUpdate,omitempty"`
	// ManagedByLabelValue is the value that is used for labeling all resources managed by the controller. The labels
	// will have key `resources.gardener.cloud/managed-by`.
	// Default: gardener
	// +optional
	ManagedByLabelValue *string `json:"managedByLabelValue,omitempty"`
}

// RootCAPublisherControllerConfig is the configuration for the root-ca-publisher controller.
type RootCAPublisherControllerConfig struct {
	// Enabled defines whether this controller is enabled.
	Enabled bool `json:"enabled"`
	// ConcurrentSyncs is the number of concurrent worker routines for this controller.
	// +optional
	ConcurrentSyncs *int `json:"concurrentSyncs,omitempty"`
	// RootCAFile is the path to a file containing the root CA.
	// +optional
	RootCAFile *string `json:"rootCAFile,omitempty"`
}

// SecretControllerConfig is the configuration for the secret controller.
type SecretControllerConfig struct {
	// ConcurrentSyncs is the number of concurrent worker routines for this controller.
	// +optional
	ConcurrentSyncs *int `json:"concurrentSyncs,omitempty"`
}

// TokenInvalidatorControllerConfig is the configuration for the token-invalidator controller.
type TokenInvalidatorControllerConfig struct {
	// Enabled defines whether this controller is enabled.
	Enabled bool `json:"enabled"`
	// ConcurrentSyncs is the number of concurrent worker routines for this controller.
	// +optional
	ConcurrentSyncs *int `json:"concurrentSyncs,omitempty"`
}

// TokenRequestorControllerConfig is the configuration for the token-requestor controller.
type TokenRequestorControllerConfig struct {
	// Enabled defines whether this controller is enabled.
	Enabled bool `json:"enabled"`
	// ConcurrentSyncs is the number of concurrent worker routines for this controller.
	// +optional
	ConcurrentSyncs *int `json:"concurrentSyncs,omitempty"`
}

// ResourceManagerWebhookConfiguration defines the configuration of the webhooks.
type ResourceManagerWebhookConfiguration struct {
	// HighAvailabilityConfig is the configuration for the high-availability-config webhook.
	HighAvailabilityConfig HighAvailabilityConfigWebhookConfig `json:"highAvailabilityConfig"`
	// PodSchedulerName is the configuration for the pod-scheduler-name webhook.
	PodSchedulerName PodSchedulerNameWebhookConfig `json:"podSchedulerName"`
	// PodTopologySpreadConstraints is the configuration for the pod-topology-spread-constraints webhook.
	PodTopologySpreadConstraints PodTopologySpreadConstraintsWebhookConfig `json:"podTopologySpreadConstraints"`
	// ProjectedTokenMount is the configuration for the projected-token-mount webhook.
	ProjectedTokenMount ProjectedTokenMountWebhookConfig `json:"projectedTokenMount"`
	// SeccompProfile is the configuration for the seccomp-profile webhook.
	SeccompProfile SeccompProfileWebhookConfig `json:"seccompProfile"`
	// TokenInvalidator is the configuration for the token-invalidator webhook.
	TokenInvalidator TokenInvalidatorWebhookConfig `json:"tokenInvalidator"`
}

// HighAvailabilityConfigWebhookConfig is the configuration for the high-availability-config webhook.
type HighAvailabilityConfigWebhookConfig struct {
	// Enabled defines whether this webhook is enabled.
	Enabled bool `json:"enabled"`
}

// PodSchedulerNameWebhookConfig is the configuration for the pod-scheduler-name webhook.
type PodSchedulerNameWebhookConfig struct {
	// Enabled defines whether this webhook is enabled.
	Enabled bool `json:"enabled"`
	// SchedulerName is the name of the scheduler that should be written into the .spec.schedulerName of pod resources.
	// +optional
	SchedulerName *string `json:"schedulerName,omitempty"`
}

// PodTopologySpreadConstraintsWebhookConfig is the configuration for the pod-topology-spread-constraints webhook.
type PodTopologySpreadConstraintsWebhookConfig struct {
	// Enabled defines whether this webhook is enabled.
	Enabled bool `json:"enabled"`
}

// ProjectedTokenMountWebhookConfig is the configuration for the projected-token-mount webhook.
type ProjectedTokenMountWebhookConfig struct {
	// Enabled defines whether this webhook is enabled.
	Enabled bool `json:"enabled"`
	// ExpirationSeconds is the number of seconds until mounted projected service account tokens expire.
	// +optional
	ExpirationSeconds *int64 `json:"expirationSeconds,omitempty"`
}

// SeccompProfileWebhookConfig is the configuration for the seccomp-profile webhook.
type SeccompProfileWebhookConfig struct {
	// Enabled defines whether this webhook is enabled.
	Enabled bool `json:"enabled"`
}

// TokenInvalidatorWebhookConfig is the configuration for the token-invalidator webhook.
type TokenInvalidatorWebhookConfig struct {
	// Enabled defines whether this webhook is enabled.
	Enabled bool `json:"enabled"`
}

const (
	// DefaultResourceClass is used as resource class if no class is specified on the command line
	DefaultResourceClass = "resources"
)
