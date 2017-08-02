package core

import (
	"github.com/Unknwon/goconfig"
	"github.com/wujunwei/vcloud/entity/resource/instance"
	"github.com/wujunwei/vcloud/entity/resource/tracker"

	"log"
	"strconv"
	"github.com/wujunwei/vcloud/pkg/factory/internalinterfaces"
)

type VmInformer interface {
	InstanceFor(filename string)
}

type vmInformer struct {
	factory internalinterfaces.InformerFactory
}

func (vi *vmInformer) InstanceFor(filename string) {
	vms := vi.GenerateVmFromFile(filename)
	vi.factory.InstanceFor(&instance.Vm{}, vms)
}

func (vi *vmInformer) GenerateVmFromFile(filename string) []*instance.Vm {
	conf, err := goconfig.LoadConfigFile(filename)
	if err != nil {
		log.Panicf("load options file error:%s", err)
	}
	cltSection, err := conf.GetSection("cluster")
	if err != nil {
		log.Panicf("read options file error:%s", err)
	}
	ctnSection, err := conf.GetSection("vm")
	if err != nil {
		log.Panicf("read options file error:%s", err)
	}
	mips, _ := strconv.ParseFloat(ctnSection["mips"], 64)
	ram, _ := strconv.ParseFloat(ctnSection["ram"], 64)
	bw, _ := strconv.ParseFloat(ctnSection["bw"], 64)
	vmNum, _ := strconv.Atoi(cltSection["vms"])
	vms := make([]*instance.Vm, 0)
	for i := 0; i < vmNum; i++ {
		mipsTracker := tracker.NewMipsTracker(mips, 0, nil)
		ramTracker := tracker.NewRamTracker(ram, 0, nil)
		bwTracker := tracker.NewBwTracker(bw, 0, nil)

		vm := instance.NewVm(i, nil, mipsTracker, ramTracker, bwTracker)
		vms = append(vms, vm)
	}

	return vms
}
