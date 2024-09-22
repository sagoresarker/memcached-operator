/*
Copyright 2024.

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

// MemcachedDeploymentSpec defines the desired state of MemcachedDeployment
type MemcachedDeploymentSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of MemcachedDeployment. Edit memcacheddeployment_types.go to remove/update
	Size  int32  `json:"size,omitempty"`
	Image string `json:"image,omitempty"`
}

// MemcachedDeploymentStatus defines the observed state of MemcachedDeployment
type MemcachedDeploymentStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Nodes []string `json:"nodes"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// MemcachedDeployment is the Schema for the memcacheddeployments API
type MemcachedDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MemcachedDeploymentSpec   `json:"spec,omitempty"`
	Status MemcachedDeploymentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MemcachedDeploymentList contains a list of MemcachedDeployment
type MemcachedDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MemcachedDeployment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MemcachedDeployment{}, &MemcachedDeploymentList{})
}
