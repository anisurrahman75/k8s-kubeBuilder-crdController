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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AppsCodeSpec defines the desired state of AppsCode
type AppsCodeSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Name     string `json:"name,omitempty"`
	Replicas *int32 `json:"replicas"`
	Image    string `json:"image"`

	Port     int32 `json:"port"`
	NodePort int32 `json:"nodePort"`
}

// AppsCodeStatus defines the observed state of AppsCode
type AppsCodeStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	AvailableReplicas int32 `json:"availableReplicas"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// AppsCode is the Schema for the appscodes API
type AppsCode struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppsCodeSpec   `json:"spec,omitempty"`
	Status AppsCodeStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// AppsCodeList contains a list of AppsCode
type AppsCodeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppsCode `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AppsCode{}, &AppsCodeList{})
}
