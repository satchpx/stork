/*
Copyright 2018 Openstorage.org

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

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/libopenstorage/stork/pkg/apis/stork/v1alpha1"
	scheme "github.com/libopenstorage/stork/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DataExportsGetter has a method to return a DataExportInterface.
// A group's client should implement this interface.
type DataExportsGetter interface {
	DataExports(namespace string) DataExportInterface
}

// DataExportInterface has methods to work with DataExport resources.
type DataExportInterface interface {
	Create(*v1alpha1.DataExport) (*v1alpha1.DataExport, error)
	Update(*v1alpha1.DataExport) (*v1alpha1.DataExport, error)
	UpdateStatus(*v1alpha1.DataExport) (*v1alpha1.DataExport, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DataExport, error)
	List(opts v1.ListOptions) (*v1alpha1.DataExportList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DataExport, err error)
	DataExportExpansion
}

// dataExports implements DataExportInterface
type dataExports struct {
	client rest.Interface
	ns     string
}

// newDataExports returns a DataExports
func newDataExports(c *StorkV1alpha1Client, namespace string) *dataExports {
	return &dataExports{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dataExport, and returns the corresponding dataExport object, and an error if there is any.
func (c *dataExports) Get(name string, options v1.GetOptions) (result *v1alpha1.DataExport, err error) {
	result = &v1alpha1.DataExport{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dataexports").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DataExports that match those selectors.
func (c *dataExports) List(opts v1.ListOptions) (result *v1alpha1.DataExportList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DataExportList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dataexports").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dataExports.
func (c *dataExports) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("dataexports").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a dataExport and creates it.  Returns the server's representation of the dataExport, and an error, if there is any.
func (c *dataExports) Create(dataExport *v1alpha1.DataExport) (result *v1alpha1.DataExport, err error) {
	result = &v1alpha1.DataExport{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("dataexports").
		Body(dataExport).
		Do().
		Into(result)
	return
}

// Update takes the representation of a dataExport and updates it. Returns the server's representation of the dataExport, and an error, if there is any.
func (c *dataExports) Update(dataExport *v1alpha1.DataExport) (result *v1alpha1.DataExport, err error) {
	result = &v1alpha1.DataExport{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dataexports").
		Name(dataExport.Name).
		Body(dataExport).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *dataExports) UpdateStatus(dataExport *v1alpha1.DataExport) (result *v1alpha1.DataExport, err error) {
	result = &v1alpha1.DataExport{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dataexports").
		Name(dataExport.Name).
		SubResource("status").
		Body(dataExport).
		Do().
		Into(result)
	return
}

// Delete takes name of the dataExport and deletes it. Returns an error if one occurs.
func (c *dataExports) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dataexports").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dataExports) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dataexports").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched dataExport.
func (c *dataExports) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DataExport, err error) {
	result = &v1alpha1.DataExport{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("dataexports").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
