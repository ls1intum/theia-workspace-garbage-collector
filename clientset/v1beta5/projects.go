package v1beta5

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"theia-workspace-garbage-collector/api/types/v1beta5"
)

type WorkspaceInterface interface {
	List(opts metav1.ListOptions) (*v1beta5.WorkspaceList, error)
	Get(name string, opts metav1.GetOptions) (*v1beta5.Workspace, error)
	Delete(name string, opts metav1.DeleteOptions) error
}

type workspaceClient struct {
	restClient rest.Interface
	ns         string
}

func (c *workspaceClient) List(opts metav1.ListOptions) (*v1beta5.WorkspaceList, error) {
	result := v1beta5.WorkspaceList{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource("workspaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(&result)

	return &result, err
}

func (c *workspaceClient) Get(name string, opts metav1.GetOptions) (*v1beta5.Workspace, error) {
	result := v1beta5.Workspace{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource("workspaces").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(&result)

	return &result, err
}

func (c *workspaceClient) Delete(name string, opts metav1.DeleteOptions) error {
	return c.restClient.
		Delete().
		Namespace(c.ns).
		Resource("workspaces").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.TODO()).
		Error()
}
