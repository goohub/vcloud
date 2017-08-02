package resource

import (
	"github.com/wujunwei/vcloud/entity/resource/instance"
	"github.com/wujunwei/vcloud/entity/resource/tracker"
)

type Host struct {
	id              int
	vmList          []instance.Vm
	mipsProvisioner *tracker.MipsProvisioner
	ramProvisioner  *tracker.RamProvisioner
	bwProvisioner   *tracker.BwProvisioner
	utilization     float64
}

func (host *Host) LaunchInstance(vm instance.Vm) {
	host.mipsProvisioner.Allocate(vm.GetId(), vm.GetMips())
	host.ramProvisioner.Allocate(vm.GetId(), vm.GetRam())
	host.bwProvisioner.Allocate(vm.GetId(), vm.GetBw())
}

func (host *Host) Claim(vm instance.Vm) bool {
	return host.mipsProvisioner.Claim(vm.GetMips()) &&
		host.ramProvisioner.Claim(vm.GetRam()) &&
		host.bwProvisioner.Claim(vm.GetBw())
}

func (host *Host) SetProvider(vmMipsProvider *tracker.MipsProvisioner, vmRamProvider *tracker.RamProvisioner, vmBwProvider *tracker.BwProvisioner) {
	host.mipsProvisioner = vmMipsProvider
	host.ramProvisioner = vmRamProvider
	host.bwProvisioner = vmBwProvider
}

func (host *Host) GetVms() []instance.Vm {
	return host.vmList
}

func (host *Host) GetId()int{
	return host.id
}

func (host *Host) SetId(id int) {
	host.id = id
}
