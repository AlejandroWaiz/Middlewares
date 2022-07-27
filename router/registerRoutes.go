package router

import (
	middlewares "github.com/AlejandroWaiz/Middlewares/middleware"
	"github.com/gin-gonic/gin"
)

func (r *Router) RegisterRoutes(s *gin.Engine) {

	s.Use(middlewares.ValidateHeaders())

	s.POST("/save", r.controller.SaveDataIntoDatabase)

	s.Run(":8080")

}
