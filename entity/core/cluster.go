package core

import (
	"github.com/wujunwei/vcloud/cmd/options"
	"github.com/wujunwei/vcloud/entity/resource"
	"github.com/wujunwei/vcloud/entity/resource/instance"
	"github.com/wujunwei/vcloud/pkg/provisioner/algorithm"
	"github.com/wujunwei/vcloud/entity/plugins/scheduler"
)

type Cluster interface {
	Run()
}

type cluster struct {
	Broker     *broker
	Datacenter *datacenter
}

func New(cfg *options.RuntimeConfig) Cluster {
	containers := cfg.ResourceFactory.PullInstance(&instance.Container{}).([]*instance.Container)
	vms := cfg.ResourceFactory.PullInstance(&instance.Vm{}).([]*instance.Vm)
	hosts := cfg.ResourceFactory.PullInstance(&resource.Host{}).([]*resource.Host)

	vmScheduler := algorithm.GetScheduler(&instance.Vm{}).(scheduler.VmScheduler)
	containerScheduler := algorithm.GetScheduler(&instance.Container{}).(scheduler.ContainerScheduler)

	broker := NewBroker(vms, containers, len(vms), len(containers))
	datacenter := NewDatacenter(hosts, nil, nil, vmScheduler, containerScheduler)

	return &cluster{
		broker,
		datacenter,
	}
}

func (clt *cluster) Run() {

	vmCap := len(clt.Broker.GetVms())
	vmReq := make(chan instance.Vm, vmCap)
	vmResp := make(chan bool, vmCap)

	containerCap := len(clt.Broker.GetContainers())
	containerReq := make(chan instance.Container, containerCap)
	containerResp := make(chan bool, containerCap)
	done := make(chan bool)

	go clt.Broker.Start(vmReq, vmResp, containerReq, containerResp, done)
	clt.Datacenter.Start(vmReq, vmResp, containerReq, containerResp, done)
}
