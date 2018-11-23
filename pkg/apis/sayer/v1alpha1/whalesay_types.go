package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// WhalesaySpec defines the desired state of Whalesay
type WhalesaySpec struct {
	Text string `json:"text"`
}

// WhalesayStatus defines the observed state of Whalesay
type WhalesayStatus struct {
	Executed bool `json:"executed"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Whalesay is the Schema for the whalesays API
// +k8s:openapi-gen=true
type Whalesay struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WhalesaySpec   `json:"spec,omitempty"`
	Status WhalesayStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WhalesayList contains a list of Whalesay
type WhalesayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Whalesay `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Whalesay{}, &WhalesayList{})
}
