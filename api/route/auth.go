package route

import (
	"gohero/api/controller"
	"gohero/bootstrap"
	"gohero/usecase"

	"github.com/gin-gonic/gin"
)

func NewAuthRoute(c *gin.RouterGroup, app *bootstrap.Application) {
	controller := &controller.AuthController{App: *app, AuthUsecase: usecase.NewAuthUsecase(app)}

	c.GET("/check-session", controller.SignIn)
	c.POST("/sign-in", controller.SignIn)
}
