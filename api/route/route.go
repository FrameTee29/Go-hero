package route

import (
	"fmt"
	"gohero/bootstrap"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RouteSetup(app *bootstrap.Application) {

	appEnv := app.Env.AppEnv

	if appEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.Use(cors.Default())

	// * Middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	router := r.Group("/api")
	// ? Health check

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Server is running ❤️ ❤️ ❤️")
	})

	// * New route each feature
	NewAuthRoute(router, app)

	port := fmt.Sprintf(":%s", app.Env.Port)

	// * Start the server
	if err := r.Run(port); err != nil {
		panic(err)
	}
}
