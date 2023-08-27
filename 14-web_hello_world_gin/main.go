package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oriastanjung/14-web_hello_world_gin/api/users"
)

func main() {
	router := gin.Default()

	// Get Handler for index

	router.GET("/", func(context *gin.Context) {
		context.IndentedJSON(http.StatusOK, "hello world!")
	})

	// buat route untuk group Users
	users.UsersRoutes(router)

	err := router.Run(":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
}
