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

package config

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	componentbaseconfig "k8s.io/component-base/config"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ResourceManagerConfiguration defines the configuration for the gardener-resource-manager.
type ResourceManagerConfiguration struct {
	metav1.TypeMeta
	// SourceClientConnection specifies the client connection settings for the proxy server
	// to use when communicating with the source apiserver.
	SourceClientConnection SourceClientConnection
	// TargetClientConnection specifies the client connection settings for the proxy server
	// to use when communicating with the target apiserver.
	TargetClientConnection *TargetClientConnection
	// LeaderElection defines the configuration of leader election client.
	LeaderElection componentbaseconfig.LeaderElectionConfiguration
	// Server defines the configuration of the HTTP server.
	Server ServerConfiguration
	// Debugging holds configuration for Debugging related features.
	Debugging *componentbaseconfig.DebuggingConfiguration
	// LogLevel is the level/severity for the logs. Must be one of [info,debug,error].
	LogLevel string
	// LogFormat is the output format for the logs. Must be one of [text,json].
	LogFormat string
	// Controllers defines the configuration of the controllers.
	Controllers ResourceManagerControllerConfiguration
	// Webhooks defines the configuration of the webhooks.
	Webhooks ResourceManagerWebhookConfiguration
}

// SourceClientConnection specifies the client connection settings
// for the proxy server to use when communicating with the seed apiserver.
type SourceClientConnection struct {
	componentbaseconfig.ClientConnectionConfiguration
	// Namespace in which the ManagedResources should be observed (defaults to "all namespaces").
	Namespace *string
	// CacheResyncPeriod specifies the duration how often the cache for the source cluster is resynced.
	CacheResyncPeriod *metav1.Duration
}

// TargetClientConnection specifies the client connection settings
// for the proxy server to use when communicating with the shoot apiserver.
type TargetClientConnection struct {
	componentbaseconfig.ClientConnectionConfiguration
	// Namespace in which controllers for the target clusters act on objects (defaults to "all namespaces").
	Namespace *string
	// DisableCachedClient specifies whether the cache for the target cluster client should be disabled. If true, then
	// each request is performed with a direct client.
	DisableCachedClient *bool
	// CacheResyncPeriod specifies the duration how often the cache for the target cluster is resynced.
	CacheResyncPeriod *metav1.Duration
}

// ServerConfiguration contains details for the HTTP(S) servers.
type ServerConfiguration struct {
	// Webhooks is the configuration for the HTTPS webhook server.
	Webhooks HTTPSServer
	// HealthProbes is the configuration for serving the healthz and readyz endpoints.
	HealthProbes *Server
	// Metrics is the configuration for serving the metrics endpoint.
	Metrics *Server
}

// Server contains information for HTTP(S) server configuration.
type Server struct {
	// BindAddress is the IP address on which to listen for the specified port.
	BindAddress string
	// Port is the port on which to serve unsecured, unauthenticated access.
	Port int
}

// HTTPSServer is the configuration for the HTTPSServer server.
type HTTPSServer struct {
	// Server is the configuration for the bind address and the port.
	Server
	// TLSServer contains information about the TLS configuration for an HTTPS server.
	TLS TLSServer
}

// TLSServer contains information about the TLS configuration for an HTTPS server.
type TLSServer struct {
	// ServerCertDir is the path to a directory containing the server's TLS certificate and key (the files must be
	// named tls.crt and tls.key respectively).
	ServerCertDir string
}

// ResourceManagerControllerConfiguration defines the configuration of the controllers.
type ResourceManagerControllerConfiguration struct {
	// ClusterID is the ID of the source cluster.
	ClusterID *string
	// ResourceClass is the name of the class in ManagedResources to filter for.
	ResourceClass *string

	// KubeletCSRApprover is the configuration for the kubelet-csr-approver controller.
	KubeletCSRApprover KubeletCSRApproverControllerConfig
	// GarbageCollector is the configuration for the garbage-collector controller.
	GarbageCollector GarbageCollectorControllerConfig
	// Health is the configuration for the health controller.
	Health HealthControllerConfig
	// ManagedResource is the configuration for the managed resource controller.
	ManagedResource ManagedResourceControllerConfig
	// RootCAPublisher is the configuration for the root-ca-publisher controller.
	RootCAPublisher RootCAPublisherControllerConfig
	// Secret is the configuration for the secret controller.
	Secret SecretControllerConfig
	// TokenInvalidator is the configuration for the token-invalidator controller.
	TokenInvalidator TokenInvalidatorControllerConfig
	// TokenRequestor is the configuration for the token-requestor controller.
	TokenRequestor TokenRequestorControllerConfig
}

// KubeletCSRApproverControllerConfig is the configuration for the kubelet-csr-approver controller.
type KubeletCSRApproverControllerConfig struct {
	// Enabled defines whether this controller is enabled.
	Enabled bool
	// ConcurrentSyncs is the number of concurrent worker routines for this controller.
	ConcurrentSyncs *int
}

// GarbageCollectorControllerConfig is the configuration for the garbage-collector controller.
type GarbageCollectorControllerConfig struct {
	// Enabled defines whether this controller is enabled.
	Enabled bool
	// SyncPeriod is the duration how often the controller performs its reconciliation.
	SyncPeriod *metav1.Duration
}

// HealthControllerConfig is the configuration for the health controller.
type HealthControllerConfig struct {
	// ConcurrentSyncs is the number of concurrent worker routines for this controller.
	ConcurrentSyncs *int
	// SyncPeriod is the duration how often the controller performs its reconciliation.
	SyncPeriod *metav1.Duration
}

// ManagedResourceControllerConfig is the configuration for the managed resource controller.
type ManagedResourceControllerConfig struct {
	// ConcurrentSyncs is the number of concurrent worker routines for this controller.
	ConcurrentSyncs *int
	// SyncPeriod is the duration how often the controller performs its reconciliation.
	SyncPeriod *metav1.Duration
	// AlwaysUpdate specifies whether resources will only be updated if their desired state differs from the actual
	// state. If true, an update request will be sent in each reconciliation independent of this condition.
	AlwaysUpdate *bool
	// ManagedByLabelValue is the value that is used for labeling all resources managed by the controller. The labels
	// will have key `resources.gardener.cloud/managed-by`.
	// Default: gardener
	ManagedByLabelValue *string
}

// RootCAPublisherControllerConfig is the configuration for the root-ca-publisher controller.
type RootCAPublisherControllerConfig struct {
	// Enabled defines whether this controller is enabled.
	Enabled bool
	// ConcurrentSyncs is the number of concurrent worker routines for this controller.
	ConcurrentSyncs *int
	// RootCAFile is the path to a file containing the root CA.
	RootCAFile *string
}

// SecretControllerConfig is the configuration for the secret controller.
type SecretControllerConfig struct {
	// ConcurrentSyncs is the number of concurrent worker routines for this controller.
	ConcurrentSyncs *int
}

// TokenInvalidatorControllerConfig is the configuration for the token-invalidator controller.
type TokenInvalidatorControllerConfig struct {
	// Enabled defines whether this controller is enabled.
	Enabled bool
	// ConcurrentSyncs is the number of concurrent worker routines for this controller.
	ConcurrentSyncs *int
}

// TokenRequestorControllerConfig is the configuration for the token-requestor controller.
type TokenRequestorControllerConfig struct {
	// Enabled defines whether this controller is enabled.
	Enabled bool
	// ConcurrentSyncs is the number of concurrent worker routines for this controller.
	ConcurrentSyncs *int
}

// ResourceManagerWebhookConfiguration defines the configuration of the webhooks.
type ResourceManagerWebhookConfiguration struct {
	// HighAvailabilityConfig is the configuration for the high-availability-config webhook.
	HighAvailabilityConfig HighAvailabilityConfigWebhookConfig
	// PodSchedulerName is the configuration for the pod-scheduler-name webhook.
	PodSchedulerName PodSchedulerNameWebhookConfig
	// PodTopologySpreadConstraints is the configuration for the pod-topology-spread-constraints webhook.
	PodTopologySpreadConstraints PodTopologySpreadConstraintsWebhookConfig
	// ProjectedTokenMount is the configuration for the projected-token-mount webhook.
	ProjectedTokenMount ProjectedTokenMountWebhookConfig
	// SeccompProfile is the configuration for the seccomp-profile webhook.
	SeccompProfile SeccompProfileWebhookConfig
	// TokenInvalidator is the configuration for the token-invalidator webhook.
	TokenInvalidator TokenInvalidatorWebhookConfig
}

// HighAvailabilityConfigWebhookConfig is the configuration for the high-availability-config webhook.
type HighAvailabilityConfigWebhookConfig struct {
	// Enabled defines whether this webhook is enabled.
	Enabled bool
}

// PodSchedulerNameWebhookConfig is the configuration for the pod-scheduler-name webhook.
type PodSchedulerNameWebhookConfig struct {
	// Enabled defines whether this webhook is enabled.
	Enabled bool
	// SchedulerName is the name of the scheduler that should be written into the .spec.schedulerName of pod resources.
	SchedulerName *string
}

// PodTopologySpreadConstraintsWebhookConfig is the configuration for the pod-topology-spread-constraints webhook.
type PodTopologySpreadConstraintsWebhookConfig struct {
	// Enabled defines whether this webhook is enabled.
	Enabled bool
}

// ProjectedTokenMountWebhookConfig is the configuration for the projected-token-mount webhook.
type ProjectedTokenMountWebhookConfig struct {
	// Enabled defines whether this webhook is enabled.
	Enabled bool
	// ExpirationSeconds is the number of seconds until mounted projected service account tokens expire.
	ExpirationSeconds *int64
}

// SeccompProfileWebhookConfig is the configuration for the seccomp-profile webhook.
type SeccompProfileWebhookConfig struct {
	// Enabled defines whether this webhook is enabled.
	Enabled bool
}

// TokenInvalidatorWebhookConfig is the configuration for the token-invalidator webhook.
type TokenInvalidatorWebhookConfig struct {
	// Enabled defines whether this webhook is enabled.
	Enabled bool
}
