package core

import (
	"github.com/Unknwon/goconfig"
	"github.com/wujunwei/vcloud/entity/resource/instance"
	"github.com/wujunwei/vcloud/entity/resource/tracker"

	"github.com/wujunwei/vcloud/pkg/factory"
	"log"
	"strconv"
)

type VmInformer interface {
	InstanceFor(filename string)
}

type vmInformer struct {
	factory factory.ResourceFactory
}

func (vi *vmInformer) InstanceFor(filename string) {
	vms := vi.GenerateVmFromFile(filename)
	vi.factory.InstanceFor(&instance.Vm{}, vms)
}

func (vi *vmInformer) GenerateVmFromFile(filename string) []*instance.Vm {
	conf, err := goconfig.LoadConfigFile(filename)
	if err != nil {
		log.Panicf("load config file error:%s", err)
	}
	cltSection, err := conf.GetSection("cluster")
	if err != nil {
		log.Panicf("read config file error:%s", err)
	}
	ctnSection, err := conf.GetSection("vm")
	if err != nil {
		log.Panicf("read config file error:%s", err)
	}
	mips, _ := strconv.ParseFloat(ctnSection["mips"], 64)
	ram, _ := strconv.ParseFloat(ctnSection["ram"], 64)
	bw, _ := strconv.ParseFloat(ctnSection["bw"], 64)
	vmNum, _ := strconv.Atoi(cltSection["vms"])
	vms := make([]*instance.Vm, 0)
	for i := 0; i < vmNum; i++ {
		mipsprovider := &tracker.MipsProvisioner{}
		mipsprovider.SetMips(mips)
		ramprovider := &tracker.RamProvisioner{}
		ramprovider.SetRam(ram)
		bwprovider := &tracker.BwProvisioner{}
		bwprovider.SetBw(bw)

		vm := &instance.Vm{}
		vm.SetId(i)
		vm.SetProvider(mipsprovider, ramprovider, bwprovider)
		vms = append(vms, vm)
	}

	return vms
}
