// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	core "github.com/gardener/gardener/pkg/apis/core"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeControllerRegistrations implements ControllerRegistrationInterface
type FakeControllerRegistrations struct {
	Fake *FakeCore
}

var controllerregistrationsResource = schema.GroupVersionResource{Group: "core.gardener.cloud", Version: "", Resource: "controllerregistrations"}

var controllerregistrationsKind = schema.GroupVersionKind{Group: "core.gardener.cloud", Version: "", Kind: "ControllerRegistration"}

// Get takes name of the controllerRegistration, and returns the corresponding controllerRegistration object, and an error if there is any.
func (c *FakeControllerRegistrations) Get(name string, options v1.GetOptions) (result *core.ControllerRegistration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(controllerregistrationsResource, name), &core.ControllerRegistration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*core.ControllerRegistration), err
}

// List takes label and field selectors, and returns the list of ControllerRegistrations that match those selectors.
func (c *FakeControllerRegistrations) List(opts v1.ListOptions) (result *core.ControllerRegistrationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(controllerregistrationsResource, controllerregistrationsKind, opts), &core.ControllerRegistrationList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &core.ControllerRegistrationList{ListMeta: obj.(*core.ControllerRegistrationList).ListMeta}
	for _, item := range obj.(*core.ControllerRegistrationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested controllerRegistrations.
func (c *FakeControllerRegistrations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(controllerregistrationsResource, opts))
}

// Create takes the representation of a controllerRegistration and creates it.  Returns the server's representation of the controllerRegistration, and an error, if there is any.
func (c *FakeControllerRegistrations) Create(controllerRegistration *core.ControllerRegistration) (result *core.ControllerRegistration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(controllerregistrationsResource, controllerRegistration), &core.ControllerRegistration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*core.ControllerRegistration), err
}

// Update takes the representation of a controllerRegistration and updates it. Returns the server's representation of the controllerRegistration, and an error, if there is any.
func (c *FakeControllerRegistrations) Update(controllerRegistration *core.ControllerRegistration) (result *core.ControllerRegistration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(controllerregistrationsResource, controllerRegistration), &core.ControllerRegistration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*core.ControllerRegistration), err
}

// Delete takes name of the controllerRegistration and deletes it. Returns an error if one occurs.
func (c *FakeControllerRegistrations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(controllerregistrationsResource, name), &core.ControllerRegistration{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeControllerRegistrations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(controllerregistrationsResource, listOptions)

	_, err := c.Fake.Invokes(action, &core.ControllerRegistrationList{})
	return err
}

// Patch applies the patch and returns the patched controllerRegistration.
func (c *FakeControllerRegistrations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.ControllerRegistration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(controllerregistrationsResource, name, pt, data, subresources...), &core.ControllerRegistration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*core.ControllerRegistration), err
}
