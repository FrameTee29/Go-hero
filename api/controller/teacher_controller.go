package controller

import (
	"gohero/bootstrap"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TeacherController struct {
	App bootstrap.Application
}

func (cl TeacherController) GetAllTeacher(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "All Teacher",
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
