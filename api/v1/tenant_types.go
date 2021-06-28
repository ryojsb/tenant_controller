/*
Copyright 2021.

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
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TenantSpec defines the desired state of Tenant
type TenantSpec struct {
	// Namespaces are the names of the namespaces that belong to the tenant
	//+kubebuiler:validation:Required
	//+kubebuiler:validation:MinItems=1
	Namespaces []string `json:"namespaces"`
	// NamespacePrefix is the prefix for the name of namespaces
	//+optional
	NamespacePrefix string `json:"namespacePrefix,omitempty"`
	// Admin is the identity with admin for the tenant
	//+kubebuiler:validation:Required
	Admin rbacv1.Subject `json:"admin"`
}

// TenantStatus defines the observed state of Tenant
type TenantStatus struct {
	//+patchMergeKey=type
	//+patchStrategy=merge
	//+listType=map
	//+listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

const (
	ConditionReady string = "Ready"
)

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:scope=Cluster
//+kubebuilder:printcolumn:name="ADMIN",type="string",JSONPath=".spec.admin.name"
//+kubebuilder:printcolumn:name="PREFIX",type="string",JSONPath=".spec.namespacePrefix"
//+kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"

// Tenant is the Schema for the tenants API
type Tenant struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TenantSpec   `json:"spec,omitempty"`
	Status TenantStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TenantList contains a list of Tenant
type TenantList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Tenant `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Tenant{}, &TenantList{})
}
