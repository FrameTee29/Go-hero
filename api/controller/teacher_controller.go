package controller

import (
	"gohero/bootstrap"
	"gohero/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TeacherController struct {
	App            bootstrap.Application
	TeacherUsecase domain.TeacherUsecase
}

func (cl TeacherController) GetAllTeacher(ctx *gin.Context) {
	message, err := cl.TeacherUsecase.GetAllTeacher()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong !!",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func (cl TeacherController) GetTeacherById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Teacher by Id",
	})
}

func (cl TeacherController) CreateTeacher(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Create Teacher",
	})
}

func (cl TeacherController) UpdateTeacher(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Update Teacher",
	})
}

func (cl TeacherController) DeleteTeacher(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Delete Teacher",
	})
}
