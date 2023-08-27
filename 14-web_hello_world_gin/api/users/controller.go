package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oriastanjung/14-web_hello_world_gin/utils/response"
)

func GetAllUsers(context *gin.Context) {
	users := []User{
		{Id: "7218hjdajk1y234", FullName: "O. Riastanjung", Email: "oriastan999@gmail.com", Password: "12345678"},
		{Id: "2118hjdajk1y234", FullName: "O. Midiyanto", Email: "omidiyanto@gmail.com", Password: "12345678"},
	}

	if len(users) == 0 {
		res := response.Response{
			Data:    "",
			Message: "No Data Found",
		}
		context.IndentedJSON(http.StatusNotFound, res)
		return
	}

	res := response.Response{
		Data:    users,
		Message: "success",
	}
	context.IndentedJSON(http.StatusOK, res)
}
