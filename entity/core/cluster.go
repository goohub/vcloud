package core

import (
	"github.com/wujunwei/vcloud/entity/plugins"
	"github.com/wujunwei/vcloud/entity/resource/instance"
	"github.com/wujunwei/vcloud/pkg/factory"
	"github.com/wujunwei/vcloud/pkg/policy"
)

type Cluster struct {
	Datacenter Datacenter
	Broker     Broker
}

func (clt *Cluster) Init() {
	clt.loadResources()
	clt.loadScheduler()
}

func (clt *Cluster) loadResources() {
	cfgpath := "github.com/wujunwei/vcloud/cmd/config/config.ini"

	resourceFactory := factory.New()

	resourceFactory.Core().Container().InstanceFor(cfgpath)
	resourceFactory.Core().Vm().InstanceFor(cfgpath)
	resourceFactory.Core().Host().InstanceFor(cfgpath)

	clt.Broker.SetContainers(resourceFactory.PullInstance(&instance.Container{}))
	clt.Broker.SetVms(resourceFactory.PullInstance(&instance.Vm{}))

	clt.Datacenter.SetHosts(hostfactory.GenerateHostFromFile(cfgpath))
}

func (clt *Cluster) loadScheduler() {
	s := &policy.Scheduler{}
	s.InitDefaultScheduler()
	sche := s.GetScheduler()
	clt.Datacenter.SetContainerScheduler(sche["containerScheduler"].(plugins.ContainerScheduler))
	clt.Datacenter.SetVmScheduler(sche["vmScheduler"].(plugins.VmScheduler))
}

func (clt *Cluster) Run() {

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
