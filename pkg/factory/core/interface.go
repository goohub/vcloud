package core

import (
	"github.com/wujunwei/vcloud/pkg/factory"
)

type Interface interface {
	Container() ContainerInformer
	Vm() VmInformer
	Host() HostInformer
}

type informer struct {
	factory factory.ResourceFactory
}

func New(factory factory.ResourceFactory) Interface {
	return &informer{factory}
}

func (i *informer) Container() ContainerInformer {
	return &containerInformer{i.factory}
}

func (i *informer) Vm() VmInformer {
	return &vmInformer{i.factory}
}

func (i *informer) Host() HostInformer {
	return &hostInformer{i.factory}
}
