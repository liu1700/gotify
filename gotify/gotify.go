package gotify

import (
	"github.com/looplab/fsm"
)

type iService interface {
	Perform(data iData) (iResponse, error)
	HandleResponse(resp iResponse) error
}

type iData interface{}
type iResponse interface{}

type gotifyFsm struct {
	serviceFsm *fsm.FSM
	service    iService
	data       iData
}

var (
	services = make(map[string]*gotifyFsm)
)

// NewService will save the service instance to a map and do fsm initialization for this service
func NewService(serviceType string, requestName string, service iService) {
	// combineKey := fmt.Sprintf("%s|%s", serviceType, requestName)
	// services[combineKey] = service
}
