package plugins

import (
	"github.com/wujunwei/vcloud/entity/resource"
	"github.com/wujunwei/vcloud/entity/resource/instance"
)

type VmScheduler interface {
	SelectHostForVm(hosts []*resource.Host, vm instance.Vm) (interface{}, bool)
	OptimizeHost(host resource.Host) instance.Vm
}

type VmSchedulerFirstFit struct{}

func (ff *VmSchedulerFirstFit) SelectHostForVm(hosts []*resource.Host, vm instance.Vm) (interface{}, bool) {
	for _, host := range hosts {
		if host.Claim(vm){
			return host, true
		}
	}
	return nil, false
}

func (ff *VmSchedulerFirstFit) OptimizeHost(host resource.Host) instance.Vm{
	vms :=host.GetVms()
	return vms[0]
}