package policy

import (
	"github.com/wujunwei/vcloud/entity/plugins"
)

type Scheduler interface {
	SchedulerFor() map[string]interface{}
}

type scheduler struct {
	schedulerInfo map[string]interface{}
}

func New() Scheduler {
	s := &scheduler{
		schedulerInfo: make(map[string]interface{}),
	}
	s.initDefaultScheduler()
	return s
}

func (sched *scheduler) initDefaultScheduler() {
	if sched.schedulerInfo == nil {
		sched.schedulerInfo = make(map[string]interface{})
	}
	sched.schedulerInfo["vmScheduler"] = &plugins.VmSchedulerFirstFit{}
	sched.schedulerInfo["containerScheduler"] = &plugins.ContainerSchedulerFirstFit{}
}

func (sched *scheduler) SchedulerFor() map[string]interface{} {
	return sched.schedulerInfo
}
