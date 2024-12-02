package v1beta5

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"theia-workspace-garbage-collector/api/types/v1beta5"
)

type V1Beta5Interface interface {
	Workspace(namespace string) WorkspaceInterface
}

type V1Beta5Client struct {
	restClient rest.Interface
}

func NewForConfig(c *rest.Config) (*V1Beta5Client, error) {
	config := *c
	config.ContentConfig.GroupVersion = &schema.GroupVersion{Group: v1beta5.GroupName, Version: v1beta5.GroupVersion}
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	config.UserAgent = rest.DefaultKubernetesUserAgent()

	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}

	return &V1Beta5Client{restClient: client}, nil
}

func (c *V1Beta5Client) Workspaces(namespace string) WorkspaceInterface {
	return &workspaceClient{
		restClient: c.restClient,
		ns:         namespace,
	}
}
