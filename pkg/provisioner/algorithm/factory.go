package algorithm

import (
	"reflect"
	"sync"
)

var (
	schedulerFactoryMutex sync.Mutex
	schedulerTable        map[reflect.Type]interface{} = make(map[reflect.Type]interface{})
)

func RegisterAlgorithm(obj interface{}, algorithm interface{}) {
	schedulerFactoryMutex.Lock()
	defer schedulerFactoryMutex.Unlock()

	key := reflect.TypeOf(obj)
	value := algorithm
	schedulerTable[key] = value
}

func GetScheduler(obj interface{}) interface{} {
	key := reflect.TypeOf(obj)
	return schedulerTable[key]
}
