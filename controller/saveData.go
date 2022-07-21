package controller

import (
	"github.com/AlejandroWaiz/Middlewares/model"
	"github.com/gin-gonic/gin"
)

func (c *Controller) SaveDataIntoDatabase(ctx *gin.Context) {

	var requestBody model.Novel

	ctx.BindJSON(&requestBody)

	err := c.service.SaveDataIntoDatabase(requestBody)

	if err != nil {

		ctx.JSON(400, gin.H{"error": err.Error()})

	}

	ctx.JSON(200, gin.H{"Result": "Data guardada."})

}
