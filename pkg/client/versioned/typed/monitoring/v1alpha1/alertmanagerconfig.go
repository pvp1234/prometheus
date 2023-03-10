// Copyright The prometheus-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1alpha1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1alpha1"
	monitoringv1alpha1 "github.com/prometheus-operator/prometheus-operator/pkg/client/applyconfiguration/monitoring/v1alpha1"
	scheme "github.com/prometheus-operator/prometheus-operator/pkg/client/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AlertmanagerConfigsGetter has a method to return a AlertmanagerConfigInterface.
// A group's client should implement this interface.
type AlertmanagerConfigsGetter interface {
	AlertmanagerConfigs(namespace string) AlertmanagerConfigInterface
}

// AlertmanagerConfigInterface has methods to work with AlertmanagerConfig resources.
type AlertmanagerConfigInterface interface {
	Create(ctx context.Context, alertmanagerConfig *v1alpha1.AlertmanagerConfig, opts v1.CreateOptions) (*v1alpha1.AlertmanagerConfig, error)
	Update(ctx context.Context, alertmanagerConfig *v1alpha1.AlertmanagerConfig, opts v1.UpdateOptions) (*v1alpha1.AlertmanagerConfig, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.AlertmanagerConfig, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.AlertmanagerConfigList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AlertmanagerConfig, err error)
	Apply(ctx context.Context, alertmanagerConfig *monitoringv1alpha1.AlertmanagerConfigApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.AlertmanagerConfig, err error)
	AlertmanagerConfigExpansion
}

// alertmanagerConfigs implements AlertmanagerConfigInterface
type alertmanagerConfigs struct {
	client rest.Interface
	ns     string
}

// newAlertmanagerConfigs returns a AlertmanagerConfigs
func newAlertmanagerConfigs(c *MonitoringV1alpha1Client, namespace string) *alertmanagerConfigs {
	return &alertmanagerConfigs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the alertmanagerConfig, and returns the corresponding alertmanagerConfig object, and an error if there is any.
func (c *alertmanagerConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AlertmanagerConfig, err error) {
	result = &v1alpha1.AlertmanagerConfig{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("alertmanagerconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AlertmanagerConfigs that match those selectors.
func (c *alertmanagerConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AlertmanagerConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AlertmanagerConfigList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("alertmanagerconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested alertmanagerConfigs.
func (c *alertmanagerConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("alertmanagerconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a alertmanagerConfig and creates it.  Returns the server's representation of the alertmanagerConfig, and an error, if there is any.
func (c *alertmanagerConfigs) Create(ctx context.Context, alertmanagerConfig *v1alpha1.AlertmanagerConfig, opts v1.CreateOptions) (result *v1alpha1.AlertmanagerConfig, err error) {
	result = &v1alpha1.AlertmanagerConfig{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("alertmanagerconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(alertmanagerConfig).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a alertmanagerConfig and updates it. Returns the server's representation of the alertmanagerConfig, and an error, if there is any.
func (c *alertmanagerConfigs) Update(ctx context.Context, alertmanagerConfig *v1alpha1.AlertmanagerConfig, opts v1.UpdateOptions) (result *v1alpha1.AlertmanagerConfig, err error) {
	result = &v1alpha1.AlertmanagerConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("alertmanagerconfigs").
		Name(alertmanagerConfig.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(alertmanagerConfig).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the alertmanagerConfig and deletes it. Returns an error if one occurs.
func (c *alertmanagerConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("alertmanagerconfigs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *alertmanagerConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("alertmanagerconfigs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched alertmanagerConfig.
func (c *alertmanagerConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AlertmanagerConfig, err error) {
	result = &v1alpha1.AlertmanagerConfig{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("alertmanagerconfigs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied alertmanagerConfig.
func (c *alertmanagerConfigs) Apply(ctx context.Context, alertmanagerConfig *monitoringv1alpha1.AlertmanagerConfigApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.AlertmanagerConfig, err error) {
	if alertmanagerConfig == nil {
		return nil, fmt.Errorf("alertmanagerConfig provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(alertmanagerConfig)
	if err != nil {
		return nil, err
	}
	name := alertmanagerConfig.Name
	if name == nil {
		return nil, fmt.Errorf("alertmanagerConfig.Name must be provided to Apply")
	}
	result = &v1alpha1.AlertmanagerConfig{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("alertmanagerconfigs").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
