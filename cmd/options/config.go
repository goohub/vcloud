package options

import (
	"github.com/wujunwei/vcloud/pkg/factory"
	"github.com/wujunwei/vcloud/pkg/policy"
)

type RuntimeConfig struct {
	ResourceFactory factory.ResourceFactory
	Scheduler       policy.Scheduler
}

func (rc *RuntimeConfig) instantiate() {
	cfgpath := "github.com/wujunwei/vcloud/cmd/options/config.ini"

	rc.ResourceFactory.Core().Container().InstanceFor(cfgpath)
	rc.ResourceFactory.Core().Vm().InstanceFor(cfgpath)
	rc.ResourceFactory.Core().Host().InstanceFor(cfgpath)
}

func NewRuntimeConfig() *RuntimeConfig {
	factory := factory.New()
	scheduler := policy.New()

	rc := &RuntimeConfig{
		factory,
		scheduler,
	}
	rc.instantiate()
	return rc
}
