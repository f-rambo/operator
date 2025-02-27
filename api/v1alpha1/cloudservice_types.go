/*
Copyright 2025.

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

// CloudServiceSpec defines the desired state of CloudService.
type CloudServiceSpec struct {
	Id            int64             `json:"id,omitempty"`
	Image         string            `json:"image,omitempty"`
	Replicas      int32             `json:"replicas,omitempty"`
	RequestCPU    int32             `json:"request_cpu,omitempty"`
	LimitCPU      int32             `json:"limit_cpu,omitempty"`
	RequestGPU    int32             `json:"request_gpu,omitempty"`
	LimitGPU      int32             `json:"limit_gpu,omitempty"`
	RequestMemory int32             `json:"request_memory,omitempty"`
	LimitMemory   int32             `json:"limit_memory,omitempty"`
	Volumes       []Volume          `json:"volumes,omitempty"`
	Ports         []Port            `json:"ports,omitempty"`
	ConfigPath    string            `json:"config_path,omitempty"`
	Config        map[string]string `json:"config,omitempty"`
	Log           string            `json:"log,omitempty"`
}

type Port struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	IngressPath   string `json:"ingress_path,omitempty"`
	Protocol      string `json:"protocol,omitempty"`
	ContainerPort int32  `json:"container_port,omitempty"`
	ServiceId     int64  `json:"service_id,omitempty"`
}

type Volume struct {
	Id        int64  `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Path      string `json:"path,omitempty"`
	Storage   int32  `json:"storage,omitempty"`
	ServiceId int64  `json:"service_id,omitempty"`
}

// CloudServiceStatus defines the observed state of CloudService.
type Status int32

const (
	StatusUnknown     Status = 0
	StatusStarting    Status = 1
	StatuseRunning    Status = 2
	StatuseTerminated Status = 3
)

type CloudServiceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Status Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// CloudService is the Schema for the cloudservices API.
type CloudService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudServiceSpec   `json:"spec,omitempty"`
	Status CloudServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudServiceList contains a list of CloudService.
type CloudServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudService{}, &CloudServiceList{})
}
