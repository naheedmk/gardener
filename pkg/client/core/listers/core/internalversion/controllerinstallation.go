// Code generated by lister-gen. DO NOT EDIT.

package internalversion

import (
	core "github.com/gardener/gardener/pkg/apis/core"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ControllerInstallationLister helps list ControllerInstallations.
type ControllerInstallationLister interface {
	// List lists all ControllerInstallations in the indexer.
	List(selector labels.Selector) (ret []*core.ControllerInstallation, err error)
	// Get retrieves the ControllerInstallation from the index for a given name.
	Get(name string) (*core.ControllerInstallation, error)
	ControllerInstallationListerExpansion
}

// controllerInstallationLister implements the ControllerInstallationLister interface.
type controllerInstallationLister struct {
	indexer cache.Indexer
}

// NewControllerInstallationLister returns a new ControllerInstallationLister.
func NewControllerInstallationLister(indexer cache.Indexer) ControllerInstallationLister {
	return &controllerInstallationLister{indexer: indexer}
}

// List lists all ControllerInstallations in the indexer.
func (s *controllerInstallationLister) List(selector labels.Selector) (ret []*core.ControllerInstallation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*core.ControllerInstallation))
	})
	return ret, err
}

// Get retrieves the ControllerInstallation from the index for a given name.
func (s *controllerInstallationLister) Get(name string) (*core.ControllerInstallation, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(core.Resource("controllerinstallation"), name)
	}
	return obj.(*core.ControllerInstallation), nil
}
