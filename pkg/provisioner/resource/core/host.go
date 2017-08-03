package core

import (
	"github.com/Unknwon/goconfig"
	"github.com/wujunwei/vcloud/entity/resource"
	"github.com/wujunwei/vcloud/entity/resource/tracker"
	"github.com/wujunwei/vcloud/pkg/provisioner/resource/internalinterfaces"

	"log"
	"strconv"
)

type HostInformer interface {
	InstanceFor(filename string)
}

type hostInformer struct {
	factory internalinterfaces.InformerFactory
}

func (hi *hostInformer) InstanceFor(filename string) {
	hosts := hi.generateHostFromFile(filename)
	hi.factory.InstanceFor(&resource.Host{}, hosts)
}

func (hi *hostInformer) generateHostFromFile(filename string) []*resource.Host {
	conf, err := goconfig.LoadConfigFile(filename)
	if err != nil {
		log.Panicf("load options file error:%s", err)
	}
	cltSection, err := conf.GetSection("cluster")
	if err != nil {
		log.Panicf("read options file error:%s", err)
	}
	ctnSection, err := conf.GetSection("host")
	if err != nil {
		log.Panicf("read options file error:%s", err)
	}
	mips, _ := strconv.ParseFloat(ctnSection["mips"], 64)
	ram, _ := strconv.ParseFloat(ctnSection["ram"], 64)
	bw, _ := strconv.ParseFloat(ctnSection["bw"], 64)
	scale, _ := strconv.Atoi(cltSection["hosts"])
	hosts := make([]*resource.Host, 0)
	for i := 0; i < scale; i++ {
		mipTracker := tracker.NewMipsTracker(mips, 0, nil)
		ramTracker := tracker.NewRamTracker(ram, 0, nil)
		bwTracker := tracker.NewBwTracker(bw, 0, nil)

		host := resource.New(i, nil, mipTracker, ramTracker, bwTracker, 0)
		hosts = append(hosts, host)
	}
	return hosts
}
