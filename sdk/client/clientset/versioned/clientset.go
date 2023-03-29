/*
Copyright The KCP Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package versioned

import (
	"fmt"
	"net/http"

	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"

	apiresourcev1alpha1 "github.com/kcp-dev/kcp/sdk/client/clientset/versioned/typed/apiresource/v1alpha1"
	apisv1alpha1 "github.com/kcp-dev/kcp/sdk/client/clientset/versioned/typed/apis/v1alpha1"
	corev1alpha1 "github.com/kcp-dev/kcp/sdk/client/clientset/versioned/typed/core/v1alpha1"
	schedulingv1alpha1 "github.com/kcp-dev/kcp/sdk/client/clientset/versioned/typed/scheduling/v1alpha1"
	tenancyv1alpha1 "github.com/kcp-dev/kcp/sdk/client/clientset/versioned/typed/tenancy/v1alpha1"
	topologyv1alpha1 "github.com/kcp-dev/kcp/sdk/client/clientset/versioned/typed/topology/v1alpha1"
	workloadv1alpha1 "github.com/kcp-dev/kcp/sdk/client/clientset/versioned/typed/workload/v1alpha1"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	ApiresourceV1alpha1() apiresourcev1alpha1.ApiresourceV1alpha1Interface
	ApisV1alpha1() apisv1alpha1.ApisV1alpha1Interface
	CoreV1alpha1() corev1alpha1.CoreV1alpha1Interface
	SchedulingV1alpha1() schedulingv1alpha1.SchedulingV1alpha1Interface
	TenancyV1alpha1() tenancyv1alpha1.TenancyV1alpha1Interface
	TopologyV1alpha1() topologyv1alpha1.TopologyV1alpha1Interface
	WorkloadV1alpha1() workloadv1alpha1.WorkloadV1alpha1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	apiresourceV1alpha1 *apiresourcev1alpha1.ApiresourceV1alpha1Client
	apisV1alpha1        *apisv1alpha1.ApisV1alpha1Client
	coreV1alpha1        *corev1alpha1.CoreV1alpha1Client
	schedulingV1alpha1  *schedulingv1alpha1.SchedulingV1alpha1Client
	tenancyV1alpha1     *tenancyv1alpha1.TenancyV1alpha1Client
	topologyV1alpha1    *topologyv1alpha1.TopologyV1alpha1Client
	workloadV1alpha1    *workloadv1alpha1.WorkloadV1alpha1Client
}

// ApiresourceV1alpha1 retrieves the ApiresourceV1alpha1Client
func (c *Clientset) ApiresourceV1alpha1() apiresourcev1alpha1.ApiresourceV1alpha1Interface {
	return c.apiresourceV1alpha1
}

// ApisV1alpha1 retrieves the ApisV1alpha1Client
func (c *Clientset) ApisV1alpha1() apisv1alpha1.ApisV1alpha1Interface {
	return c.apisV1alpha1
}

// CoreV1alpha1 retrieves the CoreV1alpha1Client
func (c *Clientset) CoreV1alpha1() corev1alpha1.CoreV1alpha1Interface {
	return c.coreV1alpha1
}

// SchedulingV1alpha1 retrieves the SchedulingV1alpha1Client
func (c *Clientset) SchedulingV1alpha1() schedulingv1alpha1.SchedulingV1alpha1Interface {
	return c.schedulingV1alpha1
}

// TenancyV1alpha1 retrieves the TenancyV1alpha1Client
func (c *Clientset) TenancyV1alpha1() tenancyv1alpha1.TenancyV1alpha1Interface {
	return c.tenancyV1alpha1
}

// TopologyV1alpha1 retrieves the TopologyV1alpha1Client
func (c *Clientset) TopologyV1alpha1() topologyv1alpha1.TopologyV1alpha1Interface {
	return c.topologyV1alpha1
}

// WorkloadV1alpha1 retrieves the WorkloadV1alpha1Client
func (c *Clientset) WorkloadV1alpha1() workloadv1alpha1.WorkloadV1alpha1Interface {
	return c.workloadV1alpha1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c

	if configShallowCopy.UserAgent == "" {
		configShallowCopy.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	// share the transport between all clients
	httpClient, err := rest.HTTPClientFor(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	return NewForConfigAndClient(&configShallowCopy, httpClient)
}

// NewForConfigAndClient creates a new Clientset for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfigAndClient will generate a rate-limiter in configShallowCopy.
func NewForConfigAndClient(c *rest.Config, httpClient *http.Client) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}

	var cs Clientset
	var err error
	cs.apiresourceV1alpha1, err = apiresourcev1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.apisV1alpha1, err = apisv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.coreV1alpha1, err = corev1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.schedulingV1alpha1, err = schedulingv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.tenancyV1alpha1, err = tenancyv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.topologyV1alpha1, err = topologyv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.workloadV1alpha1, err = workloadv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	cs, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.apiresourceV1alpha1 = apiresourcev1alpha1.New(c)
	cs.apisV1alpha1 = apisv1alpha1.New(c)
	cs.coreV1alpha1 = corev1alpha1.New(c)
	cs.schedulingV1alpha1 = schedulingv1alpha1.New(c)
	cs.tenancyV1alpha1 = tenancyv1alpha1.New(c)
	cs.topologyV1alpha1 = topologyv1alpha1.New(c)
	cs.workloadV1alpha1 = workloadv1alpha1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}