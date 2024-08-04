package initialize

import (
	"net/http"

	"github.com/GnauqTheBeast/go-ecommerce-backend-api/internal/controller"
	"github.com/GnauqTheBeast/go-ecommerce-backend-api/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.AuthenMiddleware())

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
