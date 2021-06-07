package service

import (
	"ProductApis/model"
	"ProductApis/repository"
	log "github.com/sirupsen/logrus"
)

type RouteService interface {
	GetDriverDetails(id int) model.Driver
	CreateDriver(driver model.Driver)
	GetAllDriverDetails() []model.Driver
	UpdateDriver(id int)
	DeleteDriver(id int) string
}

type routeService struct {
	repo repository.Repository
}

func NewRouteService(repo repository.Repository) RouteService {
	return &routeService{repo: repo}
}

func (r routeService) GetDriverDetails(id int) model.Driver {
	if id < 1 {
		handleError("Invalid id provided")
	}
	result, err := r.repo.GetDriverDetails(id)
	if err != nil {
		handleError(err.Error())
	}
	return result
}

func (r routeService) CreateDriver(driver model.Driver) {
	err := r.repo.CreateDriver(driver)
	if err != nil {
		handleError(err.Error())
	}
}

func (r routeService) GetAllDriverDetails() []model.Driver {
	result, err := r.repo.GetAllDriverDetails()
	if err != nil {
		handleError(err.Error())
	}
	return result
}

func (r routeService) UpdateDriver(id int) {
	panic("implement me")
}

func (r routeService) DeleteDriver(id int) string {
	err := r.repo.DeleteDriver(id)
	if err != nil {
		handleError(err.Error())
	}
	return "Deleted successfully"
}

func handleError(s string) {
	log.Error("Error occurred: {}", s)
}
