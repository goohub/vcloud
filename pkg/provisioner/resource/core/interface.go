package core

import (
	"github.com/wujunwei/vcloud/pkg/provisioner/resource/internalinterfaces"
)

type Interface interface {
	Container() ContainerInformer
	Vm() VmInformer
	Host() HostInformer
}

type informer struct {
	factory internalinterfaces.InformerFactory
}

func New(factory internalinterfaces.InformerFactory) Interface {
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
