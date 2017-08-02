package instance

import (
	"github.com/wujunwei/vcloud/entity/resource/tracker"
)

type Vm struct {
	id            int
	containerList []Container
	mipsTracker   tracker.MipsTracker
	ramTracker    tracker.RamTracker
	bwTracker     tracker.BwTracker
}

func NewVm(
	id int,
	containerList []Container,
	mipsProvisioner tracker.MipsTracker,
	ramProvisioner tracker.RamTracker,
	bwProvisioner tracker.BwTracker) *Vm {
	return &Vm{
		id,
		containerList,
		mipsProvisioner,
		ramProvisioner,
		bwProvisioner,
	}
}

func (vm *Vm) Claim(container Container) bool {
	return vm.mipsTracker.Claim(container.GetMips()) &&
		vm.ramTracker.Claim(container.GetRam()) &&
		vm.bwTracker.Claim(container.GetBw())
}

func (vm *Vm) LaunchInstance(container Container) {
	vm.mipsTracker.Allocate(container.GetId(), container.GetMips())
	vm.ramTracker.Allocate(container.GetId(), container.GetRam())
	vm.bwTracker.Allocate(container.GetId(), container.GetBw())
}

func (vm *Vm) GetMips() float64 {
	return vm.mipsTracker.GetMips()
}

func (vm *Vm) GetRam() float64 {
	return vm.ramTracker.GetRam()
}

func (vm *Vm) GetBw() float64 {
	return vm.bwTracker.GetBw()
}

func (vm *Vm) GetId() int {
	return vm.id
}
