package cmd

import(
	"github.com/wujunwei/vcloud/entity/core"
)

func Run(){
	clt := &core.Cluster{}
	clt.Init()
	clt.Run()
}

