package routers

import (
	"net/http"

	"github.com/GnauqTheBeast/go-ecommerce-backend-api/internal/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1/")
	{
		v1.GET("/ping", Pong)
		v1.GET("/pingpong", controller.NewUserController().GetUserById)
	}

	return r
}

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
