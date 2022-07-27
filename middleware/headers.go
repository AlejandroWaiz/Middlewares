package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

//ValidateHeaders validate corp headers
func ValidateHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := strings.Split(os.Getenv("Headers"), ",")

		log.Printf("headers: %v", h)

		for _, v := range h {
			if _, m := c.Request.Header[v]; m != true {
				fmt.Println(c.Request.Header)
				log.Println("Err on headers middleware")
				c.AbortWithStatus(http.StatusBadRequest)
				break
			}
		}
	}
}
