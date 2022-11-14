/*
Copyright 2022.

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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type additionalRoleBindings struct {
	RoleName  string `json:"roleName,omitempty"`
	NameSpace string `json:"nameSpace,omitempty"`
}

// TeamSpec defines the desired state of Team
type TeamSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Name of the Role for this Team object
	// +kubebuilder:validation:MaxLength=15
	// +kubebuilder:validation:MinLength=1
	RoleName string `json:"roleName,omitempty"`

	// Name of the Group for the RoleBinding
	// +kubebuilder:validation:MaxLength=15
	// +kubebuilder:validation:MinLength=1
	GroupName string `json:"groupName,omitempty"`

	// +kubebuilder:validation:MaxItems=500
	// +kubebuilder:validation:MinItems=1
	// +optional
	RoleBindings []additionalRoleBindings `json:"roleBindings,omitempty"`

	// A selector for finding namespaces for the RoleBinding
	//NameSpaceSelector string `json:"nameSpaceSelector,omitempty"`
}

// TeamStatus defines the observed state of Team
type TeamStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// A list of pointers to RoleBindings
	// +optional
	ActiveRoleBindings []corev1.ObjectReference `json:"activeRoleBindings,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Team is the Schema for the teams API
type Team struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TeamSpec   `json:"spec,omitempty"`
	Status TeamStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TeamList contains a list of Team
type TeamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Team `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Team{}, &TeamList{})
}
