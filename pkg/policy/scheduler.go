package policy

import (
	"github.com/wujunwei/vcloud/entity/plugins"
)

type Scheduler struct{
	scheduler map[string]interface{}
}

func (sched *Scheduler) InitDefaultScheduler(){
	if sched.scheduler ==nil{
		sched.scheduler = make(map[string]interface{})
	}
	sched.scheduler["vmScheduler"] = &plugins.VmSchedulerFirstFit{}
	sched.scheduler["containerScheduler"] = &plugins.ContainerSchedulerFirstFit{}
}

func (sched *Scheduler)GetScheduler()map[string]interface{}{
	return sched.scheduler
}

func (sched *Scheduler)Core(){

}
