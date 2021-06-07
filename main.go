package main

import (
	"ProductApis/controller"
	"ProductApis/model"
	"ProductApis/repository"
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
	repo := repository.NewRepository(db)
	routeService := service.NewRouteService(repo)
	routeHandler := controller.NewRouteHandler(routeService)

	registerHandlers(routeHandler)
}

func registerHandlers(controller controller.Routes) {
	router := mux.NewRouter()

	router.HandleFunc("/", controller.GetAllDriverDetails).Methods("GET")
	router.HandleFunc("/", controller.CreateDriver).Methods("POST")
	router.HandleFunc("/{id}", controller.GetDriverDetails).Methods("GET")
	router.HandleFunc("/{id}", controller.UpdateDriver).Methods("PUT")
	router.HandleFunc("/{id}", controller.DeleteDriver).Methods("DELETE")

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":8093", handler))
}

func initializeDB() *gorm.DB {
	db, err := gorm.Open("postgres", "user=user password=root dbname=gormDB sslmode=disable")

	if err != nil {
		log.Error(err.Error())
	}

	database := db.DB()

	err = database.Ping()
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&model.Driver{}, &model.Car{}, &model.Company{})
	return db
}
