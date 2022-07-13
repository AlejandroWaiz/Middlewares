package router

import (
	controller "github.com/AlejandroWaiz/Middlewares/controller"
	"github.com/gin-gonic/gin"
)

type Router struct {
	controller controller.IController
}

type IRouter interface {
	RegisterRoutes(s *gin.Engine)
}

func GetRouterInstance(c controller.IController) IRouter {

	return &Router{controller: c}

}
