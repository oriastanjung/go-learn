package users

import (
	"github.com/gin-gonic/gin"
)

func UsersRoutes(router *gin.Engine) {
	usersGroup := router.Group("/users")
	{
		usersGroup.GET("/", GetAllUsers)
		// usersGroup.PUT("/") dll
	}
}
