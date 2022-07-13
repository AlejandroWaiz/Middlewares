package controller

import (
	"github.com/AlejandroWaiz/Middlewares/service"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	service service.IService
}

type IController interface {
	SaveDataIntoDatabase(c *gin.Context)
}

func GetControllerInstance(s service.IService) IController {

	return &Controller{service: s}

}
