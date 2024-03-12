package controller

import (
	"gohero/bootstrap"
	"gohero/domain"
	"net/http"
	"strconv"

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
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		panic("Parse string to int error")
	}

	message, err := cl.TeacherUsecase.GetTeacherById(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong !!",
		})
	}

	ctx.JSON(http.StatusOK, message)
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
