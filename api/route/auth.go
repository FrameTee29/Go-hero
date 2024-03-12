package route

import (
	"gohero/api/controller"
	"gohero/bootstrap"
	"gohero/repository"
	"gohero/usecase"

	"github.com/gin-gonic/gin"
)

func NewAuthRoute(c *gin.RouterGroup, app *bootstrap.Application) {
	authRepo := repository.NewAuthRepository(app)

	controller := &controller.AuthController{App: *app, AuthUsecase: usecase.NewAuthUsecase(app, authRepo)}

	c.GET("/check-session-timeout", controller.CheckSessionTimeout)
	c.POST("/sign-in", controller.SignIn)
}
