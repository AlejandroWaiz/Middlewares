package main

import (
	"log"

	"github.com/AlejandroWaiz/Middlewares/controller"
	"github.com/AlejandroWaiz/Middlewares/firestore"
	"github.com/AlejandroWaiz/Middlewares/router"
	"github.com/AlejandroWaiz/Middlewares/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")

	database := firestore.CreateFirestoreInstance()

	service := service.GetServiceInstance(database)

	controller := controller.GetControllerInstance(service)

	router := router.GetRouterInstance(controller)

	router.RegisterRoutes(gin.Default())

	log.Println("Running")

}
