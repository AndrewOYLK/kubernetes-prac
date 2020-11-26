/*
 *	curl -v --form user=user --form password=password http://localhost:8080/login
 */

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	r := gin.Default()

	r.POST("/login", func(ctx *gin.Context) {
		var form LoginForm

		if ctx.ShouldBind(&form) == nil {
			if form.User == "user" && form.Password == "password" {
				ctx.JSON(http.StatusOK, gin.H{
					"status": "you are logged in",
				})
			} else {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"status": "unauthorized",
				})
			}
		}
	})

	r.Run(":8080")
}
