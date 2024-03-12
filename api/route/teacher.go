package route

import (
	"gohero/api/controller"
	"gohero/bootstrap"
	"gohero/repositories"
	"gohero/usecase"

	"github.com/gin-gonic/gin"
)

func NewTeacherRoute(c *gin.RouterGroup, app *bootstrap.Application) {
	teacherRepository := repositories.NewTeacherRepository(app)

	teacherUsecase := usecase.NewTeacherUsecase(app, teacherRepository)

	controller := &controller.TeacherController{App: *app, TeacherUsecase: teacherUsecase}

	c.GET("/teacher/all", controller.GetAllTeacher)
	c.GET("/teacher/:id", controller.GetTeacherById)
	c.POST("/teacher/create", controller.CreateTeacher)
	c.PATCH("/teacher/update/:id", controller.UpdateTeacher)
	c.DELETE("/teacher/delete/:id", controller.DeleteTeacher)
}
