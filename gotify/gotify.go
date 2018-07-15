package gotify

import (
	"fmt"
)

type iService interface {
	Perform(data iData) (iResponse, error)
	HandleResponse(resp iResponse) error
}

type iData interface{}
type iResponse interface{}

type gotifyFsm struct {
	// serviceFsm *fsm.FSM
	// service iService
	// data       iData
}

var (
	services = make(map[string]iService)
)

// NewService will save the service instance to a map and do fsm initialization for this service
func NewService(serviceType string, requestName string, service iService) {
	combineKey := fmt.Sprintf("%s|%s", serviceType, requestName)
	// services[combineKey] = &gotifyFsm{service: service}
	services[combineKey] = service
}

func Send(serviceType string, requestName string, data iData) error {
	combineKey := fmt.Sprintf("%s|%s", serviceType, requestName)
	service := services[combineKey]
	if service == nil {
		return fmt.Errorf("Service not found, Unknown service type: %s, or request name: %s", serviceType, requestName)
	}

	resp, err := service.Perform(data)
	if err != nil {
		return err
	}

	return service.HandleResponse(resp)
}
