package v1beta5

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// Required for the K8s API to accept the CRDs as types

// DeepCopyInto copies all properties of this object into another object of the
// same type that is provided as a pointer.
func (in *Workspace) DeepCopyInto(out *Workspace) {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	out.Spec = WorkspaceSpec{
		Name:          in.Spec.Name,
		Label:         in.Spec.Label,
		AppDefinition: in.Spec.AppDefinition,
		User:          in.Spec.User,
		Storage:       in.Spec.Storage,
		Options:       in.Spec.Options,
	}
}

// DeepCopyObject returns a generically typed copy of an object
func (in *Workspace) DeepCopyObject() runtime.Object {
	out := Workspace{}
	in.DeepCopyInto(&out)

	return &out
}

// DeepCopyObject returns a generically typed copy of an object
func (in *WorkspaceList) DeepCopyObject() runtime.Object {
	out := WorkspaceList{}
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta

	if in.Items != nil {
		out.Items = make([]Workspace, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}

	return &out
}
