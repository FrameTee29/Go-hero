package controller

import (
	"gohero/bootstrap"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CourseController struct {
	App bootstrap.Application
}

func (cl CourseController) GetAllCourse(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "All Course",
	})
}

func (cl CourseController) GetCourseById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Course by Id",
	})
}

func (cl CourseController) CreateCourse(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Create course",
	})
}

func (cl CourseController) UpdateCourse(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Update course",
	})
}

func (cl CourseController) DeleteCourse(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Delete course",
	})
}
