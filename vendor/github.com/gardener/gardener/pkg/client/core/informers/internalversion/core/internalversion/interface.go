// Code generated by informer-gen. DO NOT EDIT.

package internalversion

import (
	internalinterfaces "github.com/gardener/gardener/pkg/client/core/informers/internalversion/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ControllerInstallations returns a ControllerInstallationInformer.
	ControllerInstallations() ControllerInstallationInformer
	// ControllerRegistrations returns a ControllerRegistrationInformer.
	ControllerRegistrations() ControllerRegistrationInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// ControllerInstallations returns a ControllerInstallationInformer.
func (v *version) ControllerInstallations() ControllerInstallationInformer {
	return &controllerInstallationInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ControllerRegistrations returns a ControllerRegistrationInformer.
func (v *version) ControllerRegistrations() ControllerRegistrationInformer {
	return &controllerRegistrationInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
