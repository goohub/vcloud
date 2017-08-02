package factory

import (
	"github.com/wujunwei/vcloud/pkg/factory/core"

	"reflect"
	"sync"
	"github.com/wujunwei/vcloud/pkg/factory/internalinterfaces"
)

type ResourceFactory interface {
	internalinterfaces.InformerFactory
	Core() core.Interface
}

type resourceFactory struct {
	lock       sync.Mutex
	configInfo map[reflect.Type]interface{}
}

func New() ResourceFactory {
	return &resourceFactory{
		lock:       sync.Mutex{},
		configInfo: make(map[reflect.Type]interface{}),
	}
}

func (f *resourceFactory) Core() core.Interface {
	return core.New(f)
}

func (f *resourceFactory) InstanceFor(obj interface{}, value interface{}) {
	f.lock.Lock()
	defer f.lock.Unlock()

	key := reflect.TypeOf(obj)
	f.configInfo[key] = value
}

func (f *resourceFactory) PullInstance(obj interface{}) interface{} {
	key := reflect.TypeOf(obj)
	return f.configInfo[key]
}
