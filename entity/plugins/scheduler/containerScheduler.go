package scheduler

import (
	"github.com/wujunwei/vcloud/entity/resource/instance"
)

type ContainerScheduler interface {
	SelectVmForContainer(vms []*instance.Vm, conatiner instance.Container) (interface{}, bool)
}

type containerSchedulerFirstFit struct{}

func NewContainerScheduler() ContainerScheduler {
	return &containerSchedulerFirstFit{}
}

func (ff *containerSchedulerFirstFit) SelectVmForContainer(vms []*instance.Vm, container instance.Container) (interface{}, bool) {
	for _, vm := range vms {
		if vm.Claim(container) {
			return vm, true
		}
	}
	return nil, false
}
