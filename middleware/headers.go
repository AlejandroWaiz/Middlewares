package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

//ValidateHeaders validate corp headers
func ValidateHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := strings.Split(os.Getenv("Headers"), ",")

		for _, v := range h {
			if _, m := c.Request.Header[v]; m != true {
				fmt.Println(c.Request.Header)
				c.AbortWithStatus(http.StatusBadRequest)
				break
			}
		}
	}
}
