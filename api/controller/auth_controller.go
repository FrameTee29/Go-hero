package controller

import (
	"gohero/bootstrap"
	"gohero/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	App         bootstrap.Application
	AuthUsecase domain.AuthUsecase
}

func (cl AuthController) SignIn(ctx *gin.Context) {

	message, err := cl.AuthUsecase.SignIn()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Incorrect information",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func (cl AuthController) CheckSessionTimeout(ctx *gin.Context) {
	session, err := cl.AuthUsecase.CheckSessionTimeout()

	if err != nil {
		responseMessage := domain.ResponseMessage{}

		responseMessage.Message = "Incorrect information"

		ctx.JSON(http.StatusBadRequest, responseMessage)
	}

	ctx.JSON(http.StatusOK, session)
}
