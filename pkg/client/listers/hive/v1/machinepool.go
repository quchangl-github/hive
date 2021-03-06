// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/hive/pkg/apis/hive/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MachinePoolLister helps list MachinePools.
type MachinePoolLister interface {
	// List lists all MachinePools in the indexer.
	List(selector labels.Selector) (ret []*v1.MachinePool, err error)
	// MachinePools returns an object that can list and get MachinePools.
	MachinePools(namespace string) MachinePoolNamespaceLister
	MachinePoolListerExpansion
}

// machinePoolLister implements the MachinePoolLister interface.
type machinePoolLister struct {
	indexer cache.Indexer
}

// NewMachinePoolLister returns a new MachinePoolLister.
func NewMachinePoolLister(indexer cache.Indexer) MachinePoolLister {
	return &machinePoolLister{indexer: indexer}
}

// List lists all MachinePools in the indexer.
func (s *machinePoolLister) List(selector labels.Selector) (ret []*v1.MachinePool, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.MachinePool))
	})
	return ret, err
}

// MachinePools returns an object that can list and get MachinePools.
func (s *machinePoolLister) MachinePools(namespace string) MachinePoolNamespaceLister {
	return machinePoolNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MachinePoolNamespaceLister helps list and get MachinePools.
type MachinePoolNamespaceLister interface {
	// List lists all MachinePools in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.MachinePool, err error)
	// Get retrieves the MachinePool from the indexer for a given namespace and name.
	Get(name string) (*v1.MachinePool, error)
	MachinePoolNamespaceListerExpansion
}

// machinePoolNamespaceLister implements the MachinePoolNamespaceLister
// interface.
type machinePoolNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MachinePools in the indexer for a given namespace.
func (s machinePoolNamespaceLister) List(selector labels.Selector) (ret []*v1.MachinePool, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.MachinePool))
	})
	return ret, err
}

// Get retrieves the MachinePool from the indexer for a given namespace and name.
func (s machinePoolNamespaceLister) Get(name string) (*v1.MachinePool, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("machinepool"), name)
	}
	return obj.(*v1.MachinePool), nil
}
