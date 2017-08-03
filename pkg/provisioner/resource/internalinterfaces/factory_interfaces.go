package internalinterfaces

type InformerFactory interface {
	InstanceFor(obj interface{}, value interface{})
	PullInstance(obj interface{}) interface{}
}
