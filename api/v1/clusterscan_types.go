/*
Copyright 2024 JunbeomKim.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ConcurrencyPolicy string

const (
	// AllowConcurrent allows CronJobs to run concurrently.
	AllowConcurrent ConcurrencyPolicy = "Allow"
	// ForbidConcurrent forbids concurrent runs, skipping next run if previous
	// hasn't finished yet.
	ForbidConcurrent ConcurrencyPolicy = "Forbid"
	// ReplaceConcurrent cancels currently running job and replaces it with a new one.
	ReplaceConcurrent ConcurrencyPolicy = "Replace"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ClusterScanSpec defines the desired state of ClusterScan
type ClusterScanSpec struct {
	Schedule                   string                  `json:"schedule"`
	StartingDeadlineSeconds    *int64                  `json:"startingDeadlineSeconds,omitempty"`
	ConcurrencyPolicy          ConcurrencyPolicy       `json:"concurrencyPolicy,omitempty"`
	Suspend                    *bool                   `json:"suspend,omitempty"`
	JobTemplate                batchv1.JobTemplateSpec `json:"jobTemplate"`
	SuccessfulJobsHistoryLimit *int32                  `json:"successfulJobsHistoryLimit,omitempty"`
	FailedJobsHistoryLimit     *int32                  `json:"failedJobsHistoryLimit,omitempty"`
}

// ClusterScanStatus defines the observed state of ClusterScan
type ClusterScanStatus struct {
	Active           []corev1.ObjectReference `json:"active,omitempty"`
	LastScheduleTime *metav1.Time             `json:"lastScheduleTime,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ClusterScan is the Schema for the clusterscans API
type ClusterScan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterScanSpec   `json:"spec,omitempty"`
	Status ClusterScanStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ClusterScanList contains a list of ClusterScan
type ClusterScanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterScan `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterScan{}, &ClusterScanList{})
}
