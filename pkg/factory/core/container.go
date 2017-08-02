package core

import (
	"github.com/Unknwon/goconfig"
	"github.com/wujunwei/vcloud/entity/resource/instance"

	"log"
	"strconv"
	"github.com/wujunwei/vcloud/pkg/factory/internalinterfaces"
)

type ContainerInformer interface {
	InstanceFor(filename string)
}

type containerInformer struct {
	factory internalinterfaces.InformerFactory
}

func (ci *containerInformer) InstanceFor(filename string) {
	containers := ci.generateCtnFromFile(filename)
	ci.factory.InstanceFor(&instance.Container{}, containers)
}

func (ci *containerInformer) generateCtnFromFile(filename string) []*instance.Container {
	conf, err := goconfig.LoadConfigFile(filename)
	if err != nil {
		log.Panicf("load options file error:%s", err)
	}
	cltSection, err := conf.GetSection("cluster")
	if err != nil {
		log.Panicf("read options file error:%s", err)
	}
	ctnSection, err := conf.GetSection("container")
	if err != nil {
		log.Panicf("read options file error:%s", err)
	}
	mips, _ := strconv.ParseFloat(ctnSection["mips"], 64)
	ram, _ := strconv.ParseFloat(ctnSection["ram"], 64)
	bw, _ := strconv.ParseFloat(ctnSection["bw"], 64)
	scale, _ := strconv.Atoi(cltSection["containers"])
	containers := make([]*instance.Container, 0)
	for i := 0; i < scale; i++ {
		container := instance.NewContainer(i, mips, ram, bw)
		containers = append(containers, container)
	}
	return containers
}
