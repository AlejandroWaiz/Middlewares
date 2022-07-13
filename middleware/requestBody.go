package middlewares

import (
	"errors"
	"log"
	"net/http"

	"github.com/AlejandroWaiz/Middlewares/model"
	"github.com/gin-gonic/gin"
)

func ValidateRequestsBody() gin.HandlerFunc {

	return func(c *gin.Context) {

		var requestBody []model.Novel

		c.BindJSON(&requestBody)

		if err := validateBodysFields(c, requestBody); err != nil {

			log.Printf(("Middleware | LOG) %v"), err)

			c.AbortWithStatus(http.StatusBadRequest)

		}

	}

}

func validateBodysFields(c *gin.Context, requestBody []model.Novel) error {

	for _, novel := range requestBody {

		if novel.Author == "" {

			return errors.New("Err on Author field: 'Nil content'")

		} else if novel.Genre == "" {

			return errors.New("Err on Novel field: 'Nil content'")

		} else if novel.Name == "" {

			return errors.New("Err on Name field: 'Nil content'")

		} else if novel.Pages == 0 {

			return errors.New("Err on Pages field: 'Nil content'")

		}

	}

	return nil

}
