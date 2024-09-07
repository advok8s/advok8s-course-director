/*
Copyright 2024 Graham Dumpleton.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// WorkshopRef is a reference to a Workshop resource.
type WorkshopRef struct {
	// Name is the name of the Workshop resource.
	Name string `json:"name"`
}

// WorkshopInstanceSpec defines the desired state of WorkshopInstance
type WorkshopInstanceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// WorkshopRef is a reference to the Workshop resource.
	WorkshopRef WorkshopRef `json:"workshop"`

	// AutoRefresh indicates whether the associated workshop environment should
	// be automatically refreshed when the workshop definition is updated.
	// Defaults to false.
	AutoRefresh bool `json:"autoRefresh,omitempty"`
}

// WorkshopInstanceStatus defines the observed state of WorkshopInstance
type WorkshopInstanceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster

// WorkshopInstance is the Schema for the workshopinstances API
type WorkshopInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WorkshopInstanceSpec   `json:"spec,omitempty"`
	Status WorkshopInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkshopInstanceList contains a list of WorkshopInstance
type WorkshopInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorkshopInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&WorkshopInstance{}, &WorkshopInstanceList{})
}
