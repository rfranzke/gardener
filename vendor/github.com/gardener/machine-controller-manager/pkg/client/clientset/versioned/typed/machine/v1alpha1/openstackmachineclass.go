package v1alpha1

import (
	v1alpha1 "github.com/gardener/machine-controller-manager/pkg/apis/machine/v1alpha1"
	scheme "github.com/gardener/machine-controller-manager/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// OpenStackMachineClassesGetter has a method to return a OpenStackMachineClassInterface.
// A group's client should implement this interface.
type OpenStackMachineClassesGetter interface {
	OpenStackMachineClasses(namespace string) OpenStackMachineClassInterface
}

// OpenStackMachineClassInterface has methods to work with OpenStackMachineClass resources.
type OpenStackMachineClassInterface interface {
	Create(*v1alpha1.OpenStackMachineClass) (*v1alpha1.OpenStackMachineClass, error)
	Update(*v1alpha1.OpenStackMachineClass) (*v1alpha1.OpenStackMachineClass, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.OpenStackMachineClass, error)
	List(opts v1.ListOptions) (*v1alpha1.OpenStackMachineClassList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.OpenStackMachineClass, err error)
	OpenStackMachineClassExpansion
}

// openStackMachineClasses implements OpenStackMachineClassInterface
type openStackMachineClasses struct {
	client rest.Interface
	ns     string
}

// newOpenStackMachineClasses returns a OpenStackMachineClasses
func newOpenStackMachineClasses(c *MachineV1alpha1Client, namespace string) *openStackMachineClasses {
	return &openStackMachineClasses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the openStackMachineClass, and returns the corresponding openStackMachineClass object, and an error if there is any.
func (c *openStackMachineClasses) Get(name string, options v1.GetOptions) (result *v1alpha1.OpenStackMachineClass, err error) {
	result = &v1alpha1.OpenStackMachineClass{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("openstackmachineclasses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of OpenStackMachineClasses that match those selectors.
func (c *openStackMachineClasses) List(opts v1.ListOptions) (result *v1alpha1.OpenStackMachineClassList, err error) {
	result = &v1alpha1.OpenStackMachineClassList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("openstackmachineclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested openStackMachineClasses.
func (c *openStackMachineClasses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("openstackmachineclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a openStackMachineClass and creates it.  Returns the server's representation of the openStackMachineClass, and an error, if there is any.
func (c *openStackMachineClasses) Create(openStackMachineClass *v1alpha1.OpenStackMachineClass) (result *v1alpha1.OpenStackMachineClass, err error) {
	result = &v1alpha1.OpenStackMachineClass{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("openstackmachineclasses").
		Body(openStackMachineClass).
		Do().
		Into(result)
	return
}

// Update takes the representation of a openStackMachineClass and updates it. Returns the server's representation of the openStackMachineClass, and an error, if there is any.
func (c *openStackMachineClasses) Update(openStackMachineClass *v1alpha1.OpenStackMachineClass) (result *v1alpha1.OpenStackMachineClass, err error) {
	result = &v1alpha1.OpenStackMachineClass{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("openstackmachineclasses").
		Name(openStackMachineClass.Name).
		Body(openStackMachineClass).
		Do().
		Into(result)
	return
}

// Delete takes name of the openStackMachineClass and deletes it. Returns an error if one occurs.
func (c *openStackMachineClasses) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("openstackmachineclasses").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *openStackMachineClasses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("openstackmachineclasses").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched openStackMachineClass.
func (c *openStackMachineClasses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.OpenStackMachineClass, err error) {
	result = &v1alpha1.OpenStackMachineClass{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("openstackmachineclasses").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
