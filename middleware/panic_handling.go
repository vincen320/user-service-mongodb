package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vincen320/user-service-mongodb/model/web"
)

func PanicHandling() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			e := recover()
			if e != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, web.WebResponse{
					Code:    http.StatusInternalServerError,
					Message: "Unknown Error",
					Data:    e,
				})
			}
		}()
		c.Next()
	}
}
