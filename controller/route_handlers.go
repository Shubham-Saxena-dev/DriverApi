package controller

import (
	"ProductApis/model"
	"ProductApis/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Routes interface {
	GetDriverDetails(w http.ResponseWriter, req *http.Request)
	CreateDriver(w http.ResponseWriter, req *http.Request)
	GetAllDriverDetails(w http.ResponseWriter, req *http.Request)
	UpdateDriver(w http.ResponseWriter, req *http.Request)
	DeleteDriver(w http.ResponseWriter, req *http.Request)
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
	result := r.service.GetDriverDetails(getIdFromReq(req))
	returnResult(result, w)
}

func (r routeHandler) GetAllDriverDetails(w http.ResponseWriter, req *http.Request) {
	result := r.service.GetAllDriverDetails()
	returnResult(result, w)
}

func (r routeHandler) CreateDriver(w http.ResponseWriter, req *http.Request) {
	driver := model.Driver{}

	if err := json.NewDecoder(req.Body).Decode(&driver); err != nil {
		panic(err)
	}
	defer req.Body.Close()

	r.service.CreateDriver(driver)
}

func (r routeHandler) UpdateDriver(w http.ResponseWriter, req *http.Request) {
	r.service.UpdateDriver(getIdFromReq(req))
}

func (r routeHandler) DeleteDriver(w http.ResponseWriter, req *http.Request) {
	returnResult(r.service.DeleteDriver(getIdFromReq(req)), w)
}

func getIdFromReq(request *http.Request) int {
	vars := mux.Vars(request)
	idVar := vars["id"]
	id, err := strconv.Atoi(idVar)
	if err != nil {
		panic(err)
	}
	return id
}

func returnResult(result interface{}, w http.ResponseWriter) {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "    ")
	enc.Encode(result)
}
