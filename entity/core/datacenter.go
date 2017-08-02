package core

import (
	"github.com/wujunwei/vcloud/entity/plugins"
	"github.com/wujunwei/vcloud/entity/resource"
	"github.com/wujunwei/vcloud/entity/resource/instance"

	"log"
)

type datacenter struct {
	hostList           []*resource.Host
	vmList             []*instance.Vm
	containerList      []*instance.Container
	vmScheduler        plugins.VmScheduler
	containerScheduler plugins.ContainerScheduler
}

func NewDatacenter(
	hostList []*resource.Host,
	vmList []*instance.Vm,
	containerList []*instance.Container,
	vmScheduler plugins.VmScheduler,
	containerScheduler plugins.ContainerScheduler,
) *datacenter {
	return &datacenter{
		hostList,
		vmList,
		containerList,
		vmScheduler,
		containerScheduler,
	}
}

func (datacenter *datacenter) Start(
	vmReq chan instance.Vm, vmResp chan bool,
	containerReq chan instance.Container, containerResp chan bool, done chan bool) {

	scheduleDone := datacenter.waitInstance(vmReq, vmResp, containerReq, containerResp, done)
	<-scheduleDone
	log.Print("datacenter.next")

}

func (datacenter *datacenter) waitInstance(vmReq chan instance.Vm, vmResp chan bool,
	containerReq chan instance.Container, containerResp chan bool, done chan bool) chan bool {
	scheduleDone := make(chan bool)
	go func() {
		for {
			select {
			case vm, exist := <-vmReq:
				if exist {
					resp := datacenter.startupVm(vm)
					if resp {
						datacenter.vmList = append(datacenter.vmList, &vm)
					}
					vmResp <- resp
				}
			case container, exist := <-containerReq:
				if exist {
					resp := datacenter.buildupContainer(container)
					if resp {
						datacenter.containerList = append(datacenter.containerList, &container)
					}
					containerResp <- resp
				}
			case <-done:
				scheduleDone <- true
				break
			}
		}
	}()
	return scheduleDone
}

func (datacenter *datacenter) startupVm(vm instance.Vm) bool {
	hosts := datacenter.hostList
	host, exist := datacenter.vmScheduler.SelectHostForVm(hosts, vm)
	if !exist {
		log.Printf("no more powerful host")
		return false
	}
	target, _ := host.(*resource.Host)
	target.LaunchInstance(vm)
	return true
}

func (datacenter *datacenter) buildupContainer(container instance.Container) bool {
	vms := datacenter.vmList
	vm, exist := datacenter.containerScheduler.SelectVmForContainer(vms, container)
	if !exist {
		log.Printf("no more powerful host")
		return false
	}
	target, _ := vm.(*instance.Vm)
	target.LaunchInstance(container)
	return true
}

func (datacenter *datacenter) GetHosts() []*resource.Host {
	return datacenter.hostList
}
