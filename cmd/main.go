package main

import (
	routes "api-geolocation/pkg/routes"

	"github.com/labstack/echo/v4"
)

func main() {


	e := echo.New()

	routes.GetRoutes(e);

	e.Logger.Fatal(e.Start(":8080"))
}
