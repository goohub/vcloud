package cmd

import(
	"github.com/wujunwei/vcloud/entity/core"
	"github.com/wujunwei/vcloud/cmd/options"

	_ "github.com/wujunwei/vcloud/pkg/provisioner/algorithm"
)

func Run(){
	cfg := options.NewRuntimeConfig()

	clt :=core.New(cfg)
	clt.Run()

}

