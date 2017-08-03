package algorithm

import (
	"github.com/wujunwei/vcloud/entity/plugins/scheduler"
	"github.com/wujunwei/vcloud/entity/resource/instance"
)

func init() {
	registerVmScheduler()
	registerContainerScheduler()
}

func registerVmScheduler() {
	RegisterAlgorithm(&instance.Vm{}, scheduler.NewVmScheduler())
}

func registerContainerScheduler(){
	RegisterAlgorithm(&instance.Container{}, scheduler.NewContainerScheduler())
}