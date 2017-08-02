package plugins

import (
	"github.com/wujunwei/vcloud/entity/resource/instance"
)

type ContainerScheduler interface {
	SelectVmForContainer(vms []*instance.Vm, conatiner instance.Container) (interface{}, bool)
}

type ContainerSchedulerFirstFit struct{}

func (ff *ContainerSchedulerFirstFit) SelectVmForContainer(vms []*instance.Vm, container instance.Container) (interface{}, bool) {
	for _, vm := range vms {
		if vm.Claim(container) {
			return vm, true
		}
	}
	return nil, false
}
