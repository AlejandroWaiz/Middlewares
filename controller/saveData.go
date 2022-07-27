package controller

import (
	"encoding/json"
	"log"

	"github.com/AlejandroWaiz/Middlewares/model"
	"github.com/gin-gonic/gin"
)

func (c *Controller) SaveDataIntoDatabase(ctx *gin.Context) {

	var requestBody []model.Novel

	//err := ctx.BindJSON(&requestBody)

	err := json.NewDecoder(ctx.Request.Body).Decode(&requestBody)

	if err != nil {

		log.Printf("Err with request body: %v", err)

		ctx.JSON(400, gin.H{"error": err.Error()})

	}

	err = c.service.SaveDataIntoDatabase(requestBody)

	if err != nil {

		log.Printf("Err on controller: %v", err)

		ctx.JSON(400, gin.H{"error": err.Error()})

		return

	}

	ctx.JSON(200, gin.H{"Result": "Data guardada."})

}
