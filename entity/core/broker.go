package core

import (
	"github.com/wujunwei/vcloud/entity/resource/instance"

	"log"
)

type Broker struct {
	vmList        []*instance.Vm
	containerList []*instance.Container
	vmAck         int
	containerAck  int
}

func (broker *Broker) Start(
	vmReq chan instance.Vm, vmResp chan bool,
	containerReq chan instance.Container, containerResp chan bool, done chan bool) {

	//submit requests that launch vm
	wait := broker.allocateVms(vmReq, vmResp)
	<-wait

	//submit requests that launch container
	wait = broker.allocateContainers(containerReq, containerResp)
	<-wait

	done <- true
	log.Print("broker.next")
}

func (broker *Broker) allocateVms(vmReq chan instance.Vm, vmResp chan bool) chan bool {
	wait := make(chan bool)

	go func() {
		go broker.submitVms(vmReq)

		for {
			select {
			case <-vmResp:
				broker.vmAck--
				if broker.vmAck == 0 {
					wait <- true
					return
				}
			}
		}
	}()

	return wait
}

func (broker *Broker) submitVms(vmReq chan instance.Vm) {
	for _, vm := range broker.vmList {
		vmReq <- *vm
		log.Printf("broker.vm.sending:%d", vm.GetId())
	}
	close(vmReq)
}

func (broker *Broker) allocateContainers(containerReq chan instance.Container, containerResp chan bool) chan bool {
	wait := make(chan bool)

	go func() {
		go broker.submitContainers(containerReq)

		for {
			select {
			case <-containerResp:
				broker.containerAck--
				if broker.containerAck == 0 {
					wait <- true
					return
				}
			}
		}
	}()

	return wait
}

func (broker *Broker) submitContainers(containerReq chan instance.Container) {
	for _, container := range broker.containerList {
		containerReq <- *container
		log.Printf("broker.container.sending:%d", container.GetId())
	}
	close(containerReq)
}

func (broker *Broker) GetContainers() []*instance.Container {
	return broker.containerList
}

func (broker *Broker) SetContainers(value interface{}) {
	containers, err := value.([]*instance.Container)
	if err{
		log.Panicf("type transform err")
	}
	broker.containerList = containers
	broker.containerAck = len(containers)
}

func (broker *Broker) GetVms() []*instance.Vm {
	return broker.vmList
}

func (broker *Broker) SetVms(value interface{}) {
	vms, err := value.([]*instance.Vm)
	if err{
		log.Panicf("type transform err")
	}
	broker.vmList = vms
	broker.vmAck = len(vms)
}
