package controller

import (
	"github.com/AlejandroWaiz/Middlewares/model"
	"github.com/gin-gonic/gin"
)

func (c *Controller) SaveDataIntoDatabase(ctx *gin.Context) {

	var requestBody model.Novel

	ctx.BindJSON(&requestBody)

}
