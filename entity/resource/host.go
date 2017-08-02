package resource

import (
	"github.com/wujunwei/vcloud/entity/resource/instance"
	"github.com/wujunwei/vcloud/entity/resource/tracker"
)

type Host struct {
	id          int
	vmList      []instance.Vm
	mipsTracker tracker.MipsTracker
	ramTracker  tracker.RamTracker
	bwTracker   tracker.BwTracker
	utilization float64
}

func New(
	id int,
	vmList []instance.Vm,
	mipsProvisioner tracker.MipsTracker,
	ramProvisioner tracker.RamTracker,
	bwProvisioner tracker.BwTracker,
	utilization float64) *Host {
	return &Host{
		id,
		vmList,
		mipsProvisioner,
		ramProvisioner,
		bwProvisioner,
		utilization,
	}
}

func (host *Host) LaunchInstance(vm instance.Vm) {
	host.mipsTracker.Allocate(vm.GetId(), vm.GetMips())
	host.ramTracker.Allocate(vm.GetId(), vm.GetRam())
	host.bwTracker.Allocate(vm.GetId(), vm.GetBw())
}

func (host *Host) Claim(vm instance.Vm) bool {
	return host.mipsTracker.Claim(vm.GetMips()) &&
		host.ramTracker.Claim(vm.GetRam()) &&
		host.bwTracker.Claim(vm.GetBw())
}

func (host *Host) GetVms() []instance.Vm {
	return host.vmList
}

func (host *Host) GetId() int {
	return host.id
}
