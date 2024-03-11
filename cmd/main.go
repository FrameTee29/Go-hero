package main

import (
	"gohero/api/route"
	"gohero/bootstrap"
)

func main() {
	app := bootstrap.App()

	route.RouteSetup(&app)
}
