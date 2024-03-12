package route

import (
	"gohero/api/controller"
	"gohero/bootstrap"

	"github.com/gin-gonic/gin"
)

func NewTeacherRoute(c *gin.RouterGroup, app *bootstrap.Application) {

	controller := &controller.TeacherController{App: *app}

	c.GET("/teacher/all", controller.GetAllTeacher)
	c.GET("/teacher/:id", controller.GetTeacherById)
	c.POST("/teacher/create", controller.CreateTeacher)
	c.PATCH("/teacher/update/:id", controller.UpdateTeacher)
	c.DELETE("/teacher/delete/:id", controller.DeleteTeacher)
}
