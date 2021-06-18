package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jasonstanleyyoman/currency_be/utils"
)

func SimpleAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := os.Getenv("SIMPLE_AUTH_TOKEN")
		fmt.Println(authToken)

		if headerToken := c.Request.Header.Get("Authorization"); len(headerToken) > 0 {
			fmt.Println(authToken)
			fmt.Println(headerToken)

			if authToken == headerToken {
				c.Next()
				return
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, utils.GenerateErrorWithMessage("Wrong Authorization token"))
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, utils.GenerateErrorWithMessage("No Authorization token present in header"))
		return
	}
}
