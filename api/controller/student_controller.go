package controller

import (
	"gohero/bootstrap"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StudentController struct {
	App bootstrap.Application
}

func (cl StudentController) GetAllStudent(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "All Student",
	})
}

func (cl StudentController) GetStudentById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Student by Id",
	})
}

func (cl StudentController) CreateStudent(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Create Student",
	})
}

func (cl StudentController) UpdateStudent(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Update Student",
	})
}

func (cl StudentController) DeleteStudent(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Delete Student",
	})
}
