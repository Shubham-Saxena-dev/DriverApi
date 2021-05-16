package main

import (
	"ProductApis/controller"
	"ProductApis/model"
	"ProductApis/service"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	log.Info("Hi, this is driver car api")

	db := initializeDB()
	service := service.NewRouteService(db)
	controller := controller.NewRouteHandler(service)

	registerHandlers(controller)
}

func registerHandlers(controller controller.Routes) {
	router := mux.NewRouter()

	router.HandleFunc("/", controller.GetAllDriverDetails).Methods("GET")
	router.HandleFunc("/{id}", controller.CreateDriver).Methods("POST")
	router.HandleFunc("/", controller.GetAllDriverDetails).Methods("GET")
	router.HandleFunc("/{id}", controller.UpdateDriver).Methods("PUT")
	router.HandleFunc("/{id}", controller.DeleteDriver).Methods("DELETE")

	router.HandleFunc("/{id}", controller.AddNewCar).Methods("GET")


	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))
}

func initializeDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=db port=5432 user=postgres dbname=postgres sslmode=disable password=root")

	if err != nil {
		log.Error()
	}

	db.AutoMigrate(&model.Driver{})
	db.AutoMigrate(&model.Car{})
	db.AutoMigrate(&model.Company{})

	defer db.Close()
	return db
}
