package controller

import (
	"gohero/bootstrap"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	App bootstrap.Application
}

func (cl *AuthController) SignIn(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Sign In ❤️ ❤️ ❤️",
	})
}
