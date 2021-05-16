package controller

import (
	"ProductApis/service"
	"net/http"
)

type Routes interface {
	GetDriverDetails(w http.ResponseWriter, req *http.Request)
	CreateDriver(w http.ResponseWriter, req *http.Request)
	GetAllDriverDetails(w http.ResponseWriter, req *http.Request)
	UpdateDriver(w http.ResponseWriter, req *http.Request)
	DeleteDriver(w http.ResponseWriter, req *http.Request)
	AddNewCar(w http.ResponseWriter, req *http.Request)
}

type routeHandler struct {
	service service.RouteService
}

func NewRouteHandler(service service.RouteService) Routes {
	return &routeHandler{
		service: service,
	}
}

func (r routeHandler) GetDriverDetails(w http.ResponseWriter, req *http.Request) {
	panic("implement me")
}

func (r routeHandler) GetAllDriverDetails(w http.ResponseWriter, req *http.Request) {
	panic("implement me")
}

func (r routeHandler) CreateDriver(w http.ResponseWriter, req *http.Request) {
	panic("implement me")
}

func (r routeHandler) UpdateDriver(w http.ResponseWriter, req *http.Request) {
	panic("implement me")
}

func (r routeHandler) DeleteDriver(w http.ResponseWriter, req *http.Request) {
	panic("implement me")
}

func (r routeHandler) AddNewCar(w http.ResponseWriter, req *http.Request) {
	panic("implement me")
}
