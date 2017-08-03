package options

import (
	"github.com/wujunwei/vcloud/pkg/provisioner/resource"
)

type RuntimeConfig struct {
	ResourceFactory resource.ResourceFactory
}

func (rc *RuntimeConfig) instantiate() {
	cfgpath := "github.com/wujunwei/vcloud/cmd/options/config.ini"

	rc.ResourceFactory.Core().Container().InstanceFor(cfgpath)
	rc.ResourceFactory.Core().Vm().InstanceFor(cfgpath)
	rc.ResourceFactory.Core().Host().InstanceFor(cfgpath)
}

func NewRuntimeConfig() *RuntimeConfig {
	factory := resource.New()

	rc := &RuntimeConfig{
		factory,
	}
	rc.instantiate()
	return rc
}
