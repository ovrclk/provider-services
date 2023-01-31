/*
Copyright The Akash Network Authors.

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

package v2beta1

import (
	"context"
	"time"

	v2beta1 "github.com/akash-network/provider/pkg/apis/akash.network/v2beta1"
	scheme "github.com/akash-network/provider/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ProviderHostsGetter has a method to return a ProviderHostInterface.
// A group's client should implement this interface.
type ProviderHostsGetter interface {
	ProviderHosts(namespace string) ProviderHostInterface
}

// ProviderHostInterface has methods to work with ProviderHost resources.
type ProviderHostInterface interface {
	Create(ctx context.Context, providerHost *v2beta1.ProviderHost, opts v1.CreateOptions) (*v2beta1.ProviderHost, error)
	Update(ctx context.Context, providerHost *v2beta1.ProviderHost, opts v1.UpdateOptions) (*v2beta1.ProviderHost, error)
	UpdateStatus(ctx context.Context, providerHost *v2beta1.ProviderHost, opts v1.UpdateOptions) (*v2beta1.ProviderHost, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v2beta1.ProviderHost, error)
	List(ctx context.Context, opts v1.ListOptions) (*v2beta1.ProviderHostList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2beta1.ProviderHost, err error)
	ProviderHostExpansion
}

// providerHosts implements ProviderHostInterface
type providerHosts struct {
	client rest.Interface
	ns     string
}

// newProviderHosts returns a ProviderHosts
func newProviderHosts(c *AkashV2beta1Client, namespace string) *providerHosts {
	return &providerHosts{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the providerHost, and returns the corresponding providerHost object, and an error if there is any.
func (c *providerHosts) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2beta1.ProviderHost, err error) {
	result = &v2beta1.ProviderHost{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("providerhosts").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ProviderHosts that match those selectors.
func (c *providerHosts) List(ctx context.Context, opts v1.ListOptions) (result *v2beta1.ProviderHostList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v2beta1.ProviderHostList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("providerhosts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested providerHosts.
func (c *providerHosts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("providerhosts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a providerHost and creates it.  Returns the server's representation of the providerHost, and an error, if there is any.
func (c *providerHosts) Create(ctx context.Context, providerHost *v2beta1.ProviderHost, opts v1.CreateOptions) (result *v2beta1.ProviderHost, err error) {
	result = &v2beta1.ProviderHost{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("providerhosts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(providerHost).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a providerHost and updates it. Returns the server's representation of the providerHost, and an error, if there is any.
func (c *providerHosts) Update(ctx context.Context, providerHost *v2beta1.ProviderHost, opts v1.UpdateOptions) (result *v2beta1.ProviderHost, err error) {
	result = &v2beta1.ProviderHost{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("providerhosts").
		Name(providerHost.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(providerHost).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *providerHosts) UpdateStatus(ctx context.Context, providerHost *v2beta1.ProviderHost, opts v1.UpdateOptions) (result *v2beta1.ProviderHost, err error) {
	result = &v2beta1.ProviderHost{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("providerhosts").
		Name(providerHost.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(providerHost).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the providerHost and deletes it. Returns an error if one occurs.
func (c *providerHosts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("providerhosts").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *providerHosts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("providerhosts").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched providerHost.
func (c *providerHosts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2beta1.ProviderHost, err error) {
	result = &v2beta1.ProviderHost{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("providerhosts").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
