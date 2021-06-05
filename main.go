package main

import (
	"ProductApis/controller"
	"ProductApis/model"
	"ProductApis/service"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
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
	router.HandleFunc("/", controller.CreateDriver).Methods("POST")
	router.HandleFunc("/{id}", controller.GetDriverDetails).Methods("GET")
	router.HandleFunc("/{id}", controller.UpdateDriver).Methods("PUT")
	router.HandleFunc("/{id}", controller.DeleteDriver).Methods("DELETE")

	router.HandleFunc("/{id}", controller.AddNewCar).Methods("GET")


	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":8093", handler))
}

func initializeDB() *gorm.DB {
	// to run docker image with dabase created ==>docker run --name postgres_db --env POSTGRES_USER=user --env POSTGRES_PASSWORD=root --publish 127.0.0.1:5432:5432 --detach --restart unless-stopped postgres:13
	db, err := gorm.Open("postgres", "user=user password=root dbname=gormDB sslmode=disable")

	if err != nil {
		log.Error(err.Error())
	}

	database := db.DB()

	err = database.Ping()
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&model.Driver{})
	db.AutoMigrate(&model.Car{})
	db.AutoMigrate(&model.Company{})
	return db
}
