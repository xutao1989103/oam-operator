package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	)

type ApplicationConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ApplicationSpec `json:"spec,omitempty"`
}

type ApplicationSpec struct {
	Components []ComponentConfiguration `json:"components"`
}

type ComponentConfiguration struct {
	ComponentName string        `json:"componentName"`
	ParameterValues []NameValue `json:"parameterValues"`

}

type NameValue struct {
	Name string `json:"name"`
	Value string `json:"value"`
}


type Component struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ComponentSpec `json:"spec"`
}

type ComponentSpec struct {
	Workload Workload `json:",workload"`
	Parameters []ComponentParameter `json:"parameters"`
}

type ComponentParameter struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Description string `json:"description"`
	Require bool
	FieldPaths []string
}


type Workload struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec WorkloadSpec  `json:"spec"`
}

type WorkloadSpec struct {
	OsType string  `json:"osType"`
	Containers []Container
}

type Container struct {
	Name string
	Image string
	Env []NameValue
}