/*
Copyright 2019 caicloud authors. All rights reserved.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/caicloud/cyclone/pkg/apis/cyclone/v1alpha1"
	"github.com/caicloud/cyclone/pkg/k8s/clientset/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type CycloneV1alpha1Interface interface {
	RESTClient() rest.Interface
	ExecutionClustersGetter
	ProjectsGetter
	ResourcesGetter
	StagesGetter
	WorkflowsGetter
	WorkflowRunsGetter
	WorkflowTriggersGetter
}

// CycloneV1alpha1Client is used to interact with features provided by the cyclone.dev group.
type CycloneV1alpha1Client struct {
	restClient rest.Interface
}

func (c *CycloneV1alpha1Client) ExecutionClusters() ExecutionClusterInterface {
	return newExecutionClusters(c)
}

func (c *CycloneV1alpha1Client) Projects(namespace string) ProjectInterface {
	return newProjects(c, namespace)
}

func (c *CycloneV1alpha1Client) Resources(namespace string) ResourceInterface {
	return newResources(c, namespace)
}

func (c *CycloneV1alpha1Client) Stages(namespace string) StageInterface {
	return newStages(c, namespace)
}

func (c *CycloneV1alpha1Client) Workflows(namespace string) WorkflowInterface {
	return newWorkflows(c, namespace)
}

func (c *CycloneV1alpha1Client) WorkflowRuns(namespace string) WorkflowRunInterface {
	return newWorkflowRuns(c, namespace)
}

func (c *CycloneV1alpha1Client) WorkflowTriggers(namespace string) WorkflowTriggerInterface {
	return newWorkflowTriggers(c, namespace)
}

// NewForConfig creates a new CycloneV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*CycloneV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &CycloneV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new CycloneV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *CycloneV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new CycloneV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *CycloneV1alpha1Client {
	return &CycloneV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *CycloneV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
