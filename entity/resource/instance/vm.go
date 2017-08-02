package instance

import (
	"github.com/wujunwei/vcloud/entity/resource/tracker"
)

type Vm struct {
	id              int
	containerList   []Container
	mipsProvisioner *tracker.MipsProvisioner
	ramProvisioner  *tracker.RamProvisioner
	bwProvisioner   *tracker.BwProvisioner
}

func (vm *Vm) Claim(container Container) bool {
	return vm.mipsProvisioner.Claim(container.GetMips()) &&
		vm.ramProvisioner.Claim(container.GetRam()) &&
		vm.bwProvisioner.Claim(container.GetBw())
}

func (vm *Vm) LaunchInstance(container Container){
	vm.mipsProvisioner.Allocate(container.GetId(), container.GetMips())
	vm.ramProvisioner.Allocate(container.GetId(), container.GetRam())
	vm.bwProvisioner.Allocate(container.GetId(), container.GetBw())
}

func (vm *Vm) SetProvider(vmMipsProvider *tracker.MipsProvisioner, vmRamProvider *tracker.RamProvisioner, vmBwProvider *tracker.BwProvisioner) {
	vm.mipsProvisioner = vmMipsProvider
	vm.ramProvisioner = vmRamProvider
	vm.bwProvisioner = vmBwProvider
}

func (vm *Vm) GetMips() float64 {
	return vm.mipsProvisioner.GetMips()
}

func (vm *Vm) GetRam() float64 {
	return vm.ramProvisioner.GetRam()
}

func (vm *Vm) GetBw() float64 {
	return vm.bwProvisioner.GetBw()
}

func (vm *Vm) SetId(id int) {
	vm.id = id
}

func (vm *Vm) GetId() int {
	return vm.id
}
