package route

import (
	"gohero/api/controller"
	"gohero/bootstrap"

	"github.com/gin-gonic/gin"
)

func NewCourseRoute(c *gin.RouterGroup, app *bootstrap.Application) {

	controller := &controller.CourseController{App: *app}

	c.GET("/course/all", controller.GetAllCourse)
	c.GET("/course/:id", controller.GetCourseById)
	c.POST("/course/create", controller.CreateCourse)
	c.PATCH("/course/update/:id", controller.UpdateCourse)
	c.DELETE("/course/delete/:id", controller.DeleteCourse)
}
