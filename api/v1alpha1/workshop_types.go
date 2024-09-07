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

// DifficultyLevel is the difficulty level of the workshop
// +kubebuilder:validation:Enum=beginner;intermediate;advanced;extreme
type DifficultyLevel string

const (
	Beginner     DifficultyLevel = "beginner"
	Intermediate DifficultyLevel = "intermediate"
	Advanced     DifficultyLevel = "advanced"
	Extreme      DifficultyLevel = "extreme"
)

// WorkshopLabel is a named label to apply to the workshop
type WorkshopLabel struct {
	// Name is the name of the label
	Name string `json:"name"`

	// Value is the value of the label
	Value string `json:"value"`
}

// WorkshopSpec defines the desired state of Workshop
type WorkshopSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Title is the title of the workshop
	Title string `json:"title"`

	// Description is a description of the workshop
	Description string `json:"description"`

	// Vendor is the name of the vendor providing the workshop
	Vendor string `json:"vendor,omitempty"`

	// Authors is a list of authors of the workshop
	Authors []string `json:"authors,omitempty"`

	// Version is the version of the workshop
	Version string `json:"version,omitempty"`

	// Difficulty is the difficulty level of the workshop
	Difficulty DifficultyLevel `json:"difficulty,omitempty"`

	// Duration is the recommended duration of the workshop
	Duration metav1.Duration `json:"duration,omitempty"`

	// Labels is a list of labels to apply to the workshop
	Labels []WorkshopLabel `json:"labels,omitempty"`

	// Logo is the URL to the logo for the workshop
	Logo string `json:"logo,omitempty"`

	// Url is the URL to the workshop repository or docs
	Url string `json:"url,omitempty"`
}

// WorkshopStatus defines the observed state of Workshop
type WorkshopStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster

// Workshop is the Schema for the workshops API
type Workshop struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WorkshopSpec   `json:"spec,omitempty"`
	Status WorkshopStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkshopList contains a list of Workshop
type WorkshopList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Workshop `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Workshop{}, &WorkshopList{})
}
