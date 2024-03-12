package route

import (
	"gohero/api/controller"
	"gohero/bootstrap"

	"github.com/gin-gonic/gin"
)

func NewStudentRoute(c *gin.RouterGroup, app *bootstrap.Application) {

	controller := &controller.StudentController{App: *app}

	c.GET("/student/all", controller.GetAllStudent)
	c.GET("/student/:id", controller.GetStudentById)
	c.POST("/student/create", controller.CreateStudent)
	c.PATCH("/student/update/:id", controller.UpdateStudent)
	c.DELETE("/student/delete/:id", controller.DeleteStudent)
}
