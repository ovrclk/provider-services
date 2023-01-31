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

package fake

import (
	"context"

	v2beta2 "github.com/akash-network/provider/pkg/apis/akash.network/v2beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeInventories implements InventoryInterface
type FakeInventories struct {
	Fake *FakeAkashV2beta2
}

var inventoriesResource = schema.GroupVersionResource{Group: "akash.network", Version: "v2beta2", Resource: "inventories"}

var inventoriesKind = schema.GroupVersionKind{Group: "akash.network", Version: "v2beta2", Kind: "Inventory"}

// Get takes name of the inventory, and returns the corresponding inventory object, and an error if there is any.
func (c *FakeInventories) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2beta2.Inventory, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(inventoriesResource, name), &v2beta2.Inventory{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2beta2.Inventory), err
}

// List takes label and field selectors, and returns the list of Inventories that match those selectors.
func (c *FakeInventories) List(ctx context.Context, opts v1.ListOptions) (result *v2beta2.InventoryList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(inventoriesResource, inventoriesKind, opts), &v2beta2.InventoryList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2beta2.InventoryList{ListMeta: obj.(*v2beta2.InventoryList).ListMeta}
	for _, item := range obj.(*v2beta2.InventoryList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested inventories.
func (c *FakeInventories) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(inventoriesResource, opts))
}

// Create takes the representation of a inventory and creates it.  Returns the server's representation of the inventory, and an error, if there is any.
func (c *FakeInventories) Create(ctx context.Context, inventory *v2beta2.Inventory, opts v1.CreateOptions) (result *v2beta2.Inventory, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(inventoriesResource, inventory), &v2beta2.Inventory{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2beta2.Inventory), err
}

// Update takes the representation of a inventory and updates it. Returns the server's representation of the inventory, and an error, if there is any.
func (c *FakeInventories) Update(ctx context.Context, inventory *v2beta2.Inventory, opts v1.UpdateOptions) (result *v2beta2.Inventory, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(inventoriesResource, inventory), &v2beta2.Inventory{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2beta2.Inventory), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeInventories) UpdateStatus(ctx context.Context, inventory *v2beta2.Inventory, opts v1.UpdateOptions) (*v2beta2.Inventory, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(inventoriesResource, "status", inventory), &v2beta2.Inventory{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2beta2.Inventory), err
}

// Delete takes name of the inventory and deletes it. Returns an error if one occurs.
func (c *FakeInventories) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(inventoriesResource, name, opts), &v2beta2.Inventory{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeInventories) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(inventoriesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v2beta2.InventoryList{})
	return err
}

// Patch applies the patch and returns the patched inventory.
func (c *FakeInventories) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2beta2.Inventory, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(inventoriesResource, name, pt, data, subresources...), &v2beta2.Inventory{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2beta2.Inventory), err
}
