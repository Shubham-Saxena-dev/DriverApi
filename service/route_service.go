package service

import (
	"github.com/jinzhu/gorm"
)

type RouteService interface {
	GetDriverDetails()
	CreateDriver()
	GetAllDriverDetails()
	UpdateDriver()
	DeleteDriver()
	AddNewCar()
}

type routeService struct {
	repo *gorm.DB
}

func NewRouteService(repo *gorm.DB) RouteService {
	return &routeService{repo: repo}
}

func (r routeService) GetDriverDetails() {
	panic("implement me")
}

func (r routeService) GetAllDriverDetails() {
	panic("implement me")
}

func (r routeService) CreateDriver() {
	panic("implement me")
}

func (r routeService) UpdateDriver() {
	panic("implement me")
}

func (r routeService) DeleteDriver() {
	panic("implement me")
}

func (r routeService) AddNewCar() {
	panic("implement me")
}
