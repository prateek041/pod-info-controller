/*
Copyright 2023 prateek.io.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

const (
	// Pod Replicator (PR)

	// plural name of the pod replicator CRD
	PRPluralName = "podreplicators"

	// Kind of the pod replicator CRD
	PRKind = "PodReplicator"

	// Name of the pod replicator CRD
	PRName = PRPluralName + "." + "test.prateeksingh.tech"
)

// PodReplicatorSpec defines the desired state of PodReplicator
type PodReplicatorSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// This is the name of the custom resource
	Name string `json:"name,omitempty"`
}

// PodReplicatorStatus defines the observed state of PodReplicator
type PodReplicatorStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// This is the pod IP
	PodIP string `json:"podIp,omitempty"`
	// This is the host IP
	HostIP string `json:"hostIp,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// PodReplicator is the Schema for the podreplicators API
type PodReplicator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PodReplicatorSpec   `json:"spec,omitempty"`
	Status PodReplicatorStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PodReplicatorList contains a list of PodReplicator
type PodReplicatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PodReplicator `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PodReplicator{}, &PodReplicatorList{})
}
