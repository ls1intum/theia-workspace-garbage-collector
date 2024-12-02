package v1beta5

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// WorkspaceSpec defines the desired state of Workspace
type WorkspaceSpec struct {
	Name          string            `json:"name,omitempty"`
	Label         string            `json:"label,omitempty"`
	AppDefinition string            `json:"appDefinition,omitempty"`
	User          string            `json:"user,omitempty"`
	Storage       string            `json:"storage,omitempty"`
	Options       map[string]string `json:"options,omitempty"`
}

// WorkspaceStatus defines the observed state of Workspace
type WorkspaceStatus struct {
	OperatorStatus  string `json:"operatorStatus,omitempty"`
	OperatorMessage string `json:"operatorMessage,omitempty"`
	VolumeClaim     struct {
		Status  string `json:"status,omitempty"`
		Message string `json:"message,omitempty"`
	} `json:"volumeClaim,omitempty"`
	VolumeAttach struct {
		Status  string `json:"status,omitempty"`
		Message string `json:"message,omitempty"`
	} `json:"volumeAttach,omitempty"`
	Error string `json:"error,omitempty"`
}

// Workspace is the Schema for the Workspaces API
type Workspace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorkspaceSpec   `json:"spec,omitempty"`
	Status            WorkspaceStatus `json:"status,omitempty"`
}

// WorkspaceList contains a list of Workspace
type WorkspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Workspace `json:"items"`
}
