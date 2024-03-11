package route

import (
	"gohero/api/controller"
	"gohero/bootstrap"

	"github.com/gin-gonic/gin"
)

func NewAuthRoute(c *gin.RouterGroup, app *bootstrap.Application) {
	controller := &controller.AuthController{App: *app}

	c.POST("/sign-in", controller.SignIn)
}
